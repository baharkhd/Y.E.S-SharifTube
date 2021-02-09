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

// function* shuffle(array) {
//   var i = array.length;

//   while (i--) {
//     yield array.splice(Math.floor(Math.random() * (i + 1)), 1)[0];
//   }
// }

function Courses(props) {
  //   const courses_num = 10;
  //   const arr = Array.from(Array(courses_num), (_, index) => index + 1);
  //   var ranNums = shuffle(arr);
  //   var randomIndexes = [];

  //   console.log("arr:", arr)

  //   for (let i = 0; i < arr.length; i++) {
  //     randomIndexes.push(ranNums.next().value);
  //   }

  //   console.log("random indexes:", randomIndexes);

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
