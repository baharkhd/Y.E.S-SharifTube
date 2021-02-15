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

const COURSES_QUERY = gql`
  query GetCoursesByFilter($keyWords: [String!]!, $amount: Int!, $start: Int!) {
    coursesByKeyWords(keyWords: $keyWords, amount: $amount, start: $start) {
      id
      title
      summary
      createdAt
      prof {
        username
        name
      }
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

function JoinCourseModel({ joiningCourse, setState, username }) {
  const [courseInfo, setCourseInfo] = useState({
    courseID: "",
    token: ""
  });

  const { data, loading, error } = useQuery(COURSES_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first",
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
      console.log("add user to coure:", addUserToCourse);
      if (addUserToCourse.__typename == "Course") {
        alert("you successfully added to the class :D");
      } else {
        alert(addUserToCourse.message);
      }
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
                    color={
                      course.prof.username === username
                        ? "red"
                        : courseInfo.courseID == course.id
                        ? "blue"
                        : ""
                    }
                    onClick={() => {
                      if (course.prof.username !== username) {
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
            if (courseInfo.courseID !== "" && courseInfo.token !== "") {
              addUserToCourse();
            }
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
