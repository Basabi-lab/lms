
import React from 'react'
import AlbumCard from './album-card'

export default class AlbumCardList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { cardData: props.cardData };
  }
  add(album) {
    let data = this.state.cardData;
    data.push(album);
    this.setState({cardData: data})
  }
  render() {
    let CardNodes = this.state.cardData.map((card) => {
      return (
        <AlbumCard albumId={card.id} artistId={card.artist_id} title={card.title} genre={card.genre} year={card.year} />
      );
    });
    return (
      <div>{CardNodes}</div>
    );
  }
}

/*
AlbumCardList.propTypes = {
    cardData:React.PropTypes.arrayOf(React.PropTypes.object).isRequired
}
*/
