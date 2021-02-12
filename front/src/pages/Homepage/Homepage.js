import React from "react";
import { Card, Grid, Segment } from "semantic-ui-react";
import { Link } from "react-router-dom";

const courses = [
  {
    title: "course1",
    id: "ID1"
  },
  {
    title: "course2",
    id: "ID2"
  },
  {
    title: "course3",
    id: "ID3"
  },
  {
    title: "course4",
    id: "ID4"
  },
  {
    title: "course5",
    id: "ID5"
  },
  {
    title: "course6",
    id: "ID6"
  },
  {
    title: "course7",
    id: "ID7"
  },
  {
    title: "course8",
    id: "ID8"
  }
];

function Homepage() {
  return (
    <Segment style={{ top: 70 }}>
      <Grid columns={3}>
        {courses.map(course => {
          return (
            <Grid.Column>
              <Link to={"/course:" + course.id}>
                <Card
                  onClick={() => {
                    console.log("course id:", course.id);
                  }}
                >
                  <Card.Content>
                    <Card.Header>{course.title}</Card.Header>
                    <Card.Description>{course.id}</Card.Description>
                  </Card.Content>
                </Card>
              </Link>
            </Grid.Column>
          );
        })}
      </Grid>
    </Segment>
  );
}

export default Homepage;
