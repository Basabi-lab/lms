
import React from 'react';
import { SortablePane, Pane } from 'react-sortable-pane';


export default () => {
  const paneWidth = 240;
  const paneStyle = {
    color: "white",
    backgroundColor: "black",
    width: paneWidth,
    height: "100%",
    margin: 0,
    padding: 0
  };
  return (
    <div style={{width:"100vw",height:"100vh"}}>
      <SortablePane
        direction="horizontal"
        margin={20}
      >
        <Pane id={0} key={0} style={paneStyle}>
          <p>0</p>
        </Pane>
        <Pane id={1} key={1} style={paneStyle}>
          <p>1</p>
        </Pane>
      </SortablePane>
    </div>
  );
};
