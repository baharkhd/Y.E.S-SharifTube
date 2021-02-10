import React, { useState } from "react";
import {
  Modal,
  Button,
  Form,
  Label,
  Input,
  TextArea,
  Grid,
  Card
} from "semantic-ui-react";

const otherClasses = [
  {
    title: "class1",
    summary: "summary1",
    id: "ID1"
  },
  {
    title: "class2",
    summary: "summary2",
    id: "ID2"
  },
  {
    title: "class3",
    summary: "summary3",
    id: "ID3"
  },
  {
    title: "class4",
    summary: "summary4",
    id: "ID4"
  },
  {
    title: "class5",
    summary: "summary5",
    id: "ID5"
  },
  {
    title: "class6",
    summary: "summary6",
    id: "ID6"
  },
  {
    title: "class7",
    summary: "summary7",
    id: "ID7"
  },
  {
    title: "class8",
    summary: "summary8",
    id: "ID8"
  }
];

function JoinCourseModel({ joiningCourse, setState }) {
  const [newCourses, setNewCourses] = useState(
    new Array(otherClasses.length).fill(0)
  );

  return (
    <Modal open={joiningCourse}>
      <Modal.Header>Join other classes</Modal.Header>
      <Modal.Content>
        <Grid columns={2}>
          {otherClasses.map((course, i) => {
            console.log("here ", i);
            return (
              <Grid.Column>
                <Card
                  color={newCourses[i] ? "blue" : ""}
                  onClick={() => {
                    let newArray = [...newCourses];
                    newArray[i] = Math.abs(newArray[i] - 1);
                    setNewCourses(newArray);
                  }}
                >
                  <Card.Content>
                    <Card.Header>{course.title}</Card.Header>
                    <Card.Description>{course.summary}</Card.Description>
                    <Card.Meta>{course.id}</Card.Meta>
                  </Card.Content>
                </Card>
              </Grid.Column>
            );
          })}
        </Grid>
      </Modal.Content>
      <Modal.Actions>
        <Button
          positive
          onClick={() => {
            // Join class
            setState({ joiningCourse: false });
            alert(newCourses)
          }}
        >
          Join
        </Button>
        <Button
          negative
          onClick={() => {
            setState({ joiningCourse: false });
          }}
        >
          {" "}
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
}

export default JoinCourseModel;
