import React, {useState, useEffect, useCallback, useRef} from 'react';
import axios from 'axios'

const Grid = ({data, allowEdits, setData}) => {
  const [height, setHeight] = useState(100)

  const self = useCallback(node => {
    if (node !== null) {
      setHeight(node.getBoundingClientRect().width);
    }
  }, []);

  const toggle = useCallback((row,col) => {
    var d = []
    data.forEach((row, r) => {
      d[r] = Array.from(row)
    })

    d[row][col] = !data[row][col]
    setData(d)
  }, [data, setData])

  console.log("grid", data)

  if(!data) {
    return "No data"
  }

  return (
    <div ref={self} style={{
      display: "grid",
      width: "100%",
      height: height,
      background: "#eee",
      gridTemplateColumns: "1fr 1fr 1fr 1fr 1fr",
      gridAutoRows: "1fr",
      gridGap: 1,
    }}>
    {
      data.flatMap( (row, r) => (row.map((_, c) => <div
        onClick={() => allowEdits && toggle(r,c)}
        style={{
          backgroundColor: data[r][c] ? "red" : "white",
          cursor: allowEdits ? "pointer" : "cursor",
        }}></div>)) )
    }</div>
  )
}

const Player = ({data: initialData}) => {
  const [data, setData] = useState(initialData)

  useInterval(() => {
    async function fetchData() {
      try {
        console.log("send", data)
        const result = await axios({
          method: 'post',
          url: '/next',
          data: data,
        });

        setData(result.data)
      } catch(err) {
        console.log(err)
        return
      }
    }

    fetchData()
  }, 1000)

  return <Grid data={data} />
}

function App() {
  const [editing, setEditing] = useState(true)
  const [data, setData] = useState([[false, false, false, false, false],[false, false, true,true,false],[false,true,true,true,false],[false,true,false,false,false],[false,true,true,false,false],[false,false,false,false,false]])

  return (
    <>
    <div style={{display: "flex", justifyContent: "center"}}>
      {editing ?
         <div style={{cursor: "pointer"}} onClick={() => setEditing(false)}>Play</div> :
         <div style={{cursor: "pointer"}} onClick={() => setEditing(true)}>Edit Starting Position</div> 
      }
    </div>
    { editing ? <Grid data={data} setData={setData} allowEdits /> : <Player data={data} /> }
    </>
  );
}

// from https://overreacted.io/making-setinterval-declarative-with-react-hooks/
function useInterval(callback, delay) {
  const savedCallback = useRef();

  // Remember the latest callback.
  useEffect(() => {
    savedCallback.current = callback;
  }, [callback]);

  // Set up the interval.
  useEffect(() => {
    function tick() {
      savedCallback.current();
    }
    if (delay !== null) {
      let id = setInterval(tick, delay);
      return () => clearInterval(id);
    }
  }, [delay]);
}

export default App;
