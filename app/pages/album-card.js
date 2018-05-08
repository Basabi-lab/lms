
import React from 'react'

const cardStyle = {
  backgroundColor: "gray",
  color: "white",
  width: 240,
  height: 180,
  border: "1px solid black"
}

export default class AlbumCard extends React.Component {
  render() {
    return (
      <div album-id={this.props.albumId} artist-id={this.props.artistId} key={this.props.key} style={cardStyle}>
        <p>{this.props.title}</p>
        <p>{this.props.genre}</p>
        <p>{this.props.year}</p>
      </div>
    )
  }
}

/*
AlbumCard.propTypes = {
  title: React.PropTypes.string.isRequired,
  genre: React.PropTypes.string,
  year: React.PropTypes.number
}
*/
