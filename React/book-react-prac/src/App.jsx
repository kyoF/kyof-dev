import { useState, memo, useCallback, useContext } from "react";
import { Child1 } from "./components/Child1";
import { Child4 } from "./components/Child4";
import { Card } from "./components/Card";
import { AdminFlagContext } from "./components/providers/AdminFlagProvider";

export const App = memo(() => {
  console.log("App rendering");
  const [num, setNum] = useState(0);
  const onClickButton = () => {
    setNum(num+1);
  };
  const onClickReset = useCallback(() => {
    setNum(0);
  }, []);

  const { isAdmin, setIsAdmin } = useContext(AdminFlagContext)
  const onClickSwitch = () => setIsAdmin(!isAdmin);

  return (
    <>
      {isAdmin ? <span>Yes! Admin!</span> : <span>NO... Admin...</span>}
      <button onClick={onClickSwitch}>switch</button>
      <Card />
      <hr />
      <button onClick={onClickButton}>button</button>
      <p>{num}</p>
      <Child1 onClickReset={onClickReset}/>
      <Child4 />
    </>
  )
})