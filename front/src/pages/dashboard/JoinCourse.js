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
import { useMutation, gql, useQuery } from "@apollo/client";

// const otherClasses = [
//   {
//     title: "class1",
//     summary: "summary1",
//     id: "ID1"
//   },
//   {
//     title: "class2",
//     summary: "summary2",
//     id: "ID2"
//   },
//   {
//     title: "class3",
//     summary: "summary3",
//     id: "ID3"
//   },
//   {
//     title: "class4",
//     summary: "summary4",
//     id: "ID4"
//   },
//   {
//     title: "class5",
//     summary: "summary5",
//     id: "ID5"
//   },
//   {
//     title: "class6",
//     summary: "summary6",
//     id: "ID6"
//   },
//   {
//     title: "class7",
//     summary: "summary7",
//     id: "ID7"
//   },
//   {
//     title: "class8",
//     summary: "summary8",
//     id: "ID8"
//   }
// ];

const COURSES_QUERY = gql`
  query GetCoursesByFilter($keyWords: [String!]!, $amount: Int!, $start: Int!) {
    coursesByKeyWords(keyWords: $keyWords, amount: $amount, start: $start) {
      id
      title
      summary
      createdAt
      token
      # prof {
      #   username
      #   name
      # }
    }
  }
`;

const JOIN_COURSE_MUTATION = gql`
  mutation JoinCourse($courseID: String!, $token: String!) {
    addUserToCourse(courseID: $courseID, token: $token) {
      __typename
      ... on Course {
        id
        title
        summary
        createdAt
      }
      ... on Exception {
        message
      }
    }
  }
`;

function JoinCourseModel({ joiningCourse, setState }) {
  const [newCourses, setNewCourses] = useState(new Array(10).fill(0));

  const [courseInfo, setCourseInfo] = useState({
    courseID: "",
    token: ""
  });

  const { data, loading, error } = useQuery(COURSES_QUERY, {
    variables: {
      keyWords: [],
      amount: 100,
      start: 0
    },
    onError(err) {
      console.log("error in getCourses:", err);
    }
  });

  const [addUserToCourse] = useMutation(JOIN_COURSE_MUTATION, {
    variables: {
      courseID: courseInfo.courseID,
      token: courseInfo.token
    },
    onCompleted: ({ addUserToCourse }) => {
      console.log("add user to course completed:", addUserToCourse);
    }
  });

  return (
    <Modal open={joiningCourse}>
      <Modal.Header>Join other classes</Modal.Header>
      <Modal.Content scrolling>
        <Grid columns={2} stackable>
          {!loading &&
            data.coursesByKeyWords.map((course, i) => {
              let date = new Date(course.createdAt * 1000).toISOString();
              // .substr(11, 8);

              return (
                <Grid.Column>
                  <Card
                    fluid
                    color={courseInfo.courseID == course.id ? "blue" : ""}
                    onClick={() => {
                      if (courseInfo.courseID === course.id) {
                        setCourseInfo({
                          ...courseInfo,
                          courseID: ""
                        });
                      } else {
                        setCourseInfo({
                          ...courseInfo,
                          courseID: course.id
                        });
                      }
                    }}
                  >
                    <Card.Content>
                      <Card.Header>{course.title}</Card.Header>
                      <Card.Description>{course.summary}</Card.Description>
                      <Card.Meta>Created At {date}</Card.Meta>
                    </Card.Content>
                  </Card>
                  {courseInfo.courseID === course.id && (
                    <Input
                      fluid
                      placeholder={"password for class " + course.title}
                      onChange={e => {
                        setCourseInfo({ ...courseInfo, token: e.target.value });
                      }}
                    />
                  )}
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
            addUserToCourse();
            setState({ joiningCourse: false });
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
