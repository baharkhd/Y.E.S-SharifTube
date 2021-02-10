import React from "react";
import { Grid, Segment, Image, Placeholder } from "semantic-ui-react";

const avatars = [
  "https://semantic-ui.com/images/avatar/small/chris.jpg",
  "https://semantic-ui.com/images/avatar/small/ade.jpg",
  "https://semantic-ui.com/images/avatar/small/christian.jpg",
  "https://semantic-ui.com/images/avatar/small/daniel.jpg",
  "https://semantic-ui.com/images/avatar/small/elliot.jpg",
  "https://semantic-ui.com/images/avatar/small/helen.jpg",
  "https://semantic-ui.com/images/avatar/small/jenny.jpg",
  "https://semantic-ui.com/images/avatar/small/joe.jpg",
  "https://semantic-ui.com/images/avatar/small/justen.jpg",
  "https://semantic-ui.com/images/avatar/small/laura.jpg",
  "https://semantic-ui.com/images/avatar/small/matt.jpg",
  "https://semantic-ui.com/images/avatar/small/nan.jpg",
  "https://semantic-ui.com/images/avatar/small/steve.jpg",
  "https://semantic-ui.com/images/avatar/small/stevie.jpg",
  "https://semantic-ui.com/images/avatar/small/veronika.jpg"
];


function Courses(props) {

  return (
    <Segment
      style={{
        position: "absolute",
        left: props.isMobile ? 0 : 250,
        right: 0,
        margin: 30,
        top: 70
      }}
    >
      <Grid columns={3}>
        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>

        <Grid.Column>
          <Placeholder>
            <Placeholder.Image rectangular />
          </Placeholder>
        </Grid.Column>
      </Grid>
    </Segment>
  );
}

export default Courses;
