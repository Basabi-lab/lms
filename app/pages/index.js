
import React from 'react';
import { SortablePane, Pane } from 'react-sortable-pane';
import AlbumCardList from './album-card-list'

const paneWidth = 240;
const paneStyle = {
  color: "white",
  backgroundColor: "black",
  width: paneWidth,
  height: "100%"
};

const data = [
  {
    id: 0,
    title: "Hello world",
    artist_id: 0,
    genre: "Rock",
    year: 2014
  },
  {
    id: 1,
    title: "Overflow",
    artist_id: 1,
    genre: "Rock",
    year: 2013
  }
];


export default () => {
  return (
    <div style={{width:"100vw",height:"100vh"}}>
      <SortablePane direction="horizontal" margin={20}>
        <Pane id={0} key={0} width={240} height="100%" style={paneStyle}>
          <p>Albums</p>
          <AlbumCardList cardData={data} />
        </Pane>
        <Pane id={1} key={1} width={240} height="100%" style={paneStyle}>
          <p>Songs</p>
        </Pane>
      </SortablePane>
    </div>
  );
};
