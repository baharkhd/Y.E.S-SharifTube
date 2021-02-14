import React, { useState } from "react";
import {
  Grid,
  Segment,
  Image,
  Placeholder,
  Card,
  Divider,
  Header,
  Icon
} from "semantic-ui-react";
import { Link } from "react-router-dom";
import { gql, useQuery } from "@apollo/client";

const GET_USER_QUERY = gql`
  {
    user {
      username
      name
      email
      courseIDs
    }
  }
`;

const GET_COURSES_QUERY = gql`
  query GetCourses($ids: [String!]!) {
    courses(ids: $ids) {
      id
      title
      summary
      prof {
        username
        name
        email
      }
      tas {
        username
        name
        email
      }
    }
  }
`;

function Courses(props) {
  const [courseIDs, setCourseIDs] = useState([]);

  const courses = useQuery(GET_COURSES_QUERY, {
    variables: {
      ids: courseIDs
    }
  });

  const { data, loading, error } = useQuery(GET_USER_QUERY, {
    onCompleted: ({ user }) => {
      console.log("user:", user);
      setCourseIDs(user.courseIDs);
    }
  });

  let test = [1, 2, 3, 4];
  test = test.filter(t => {
    return t == 2;
  });
  console.log("test:", test);

  console.log("coursesObject:", courses);
  console.log("coursesIDs", courseIDs);
  console.log("user courses:", data);
  console.log("loading:", loading);
  console.log("errror:", error);

  let yourClasses, otherClasses;

  if (courses.data) {
    yourClasses = courses.data.courses.filter(c => {
      return c.prof.username == data.user.username;
    });
    otherClasses = courses.data.courses.filter(c => {
      return c.prof.username != data.user.username;
    });
  }

  // let yourClasses = courses.data.courses.filter(c => {
  //   return c.prof.username == data.user.username;
  // });

  // let otherClasses = courses.data.courses.filter(c => {
  //   return c.prof.username != data.user.username;
  // });

  return (
    <Segment
      style={{
        position: "absolute",
        left: props.isMobile ? 0 : 250,
        right: 0,
        margin: 30,
        top: 70,
        padding: 10
      }}
    >
      <Divider horizontal>
        <Header textAlign="left">
          <Icon name="video play" />
          Your Courses
        </Header>
      </Divider>
      <Grid columns={2} stackable>
        {!courses.loading &&
          yourClasses.map(course => {
            // let date = new Date(course.createdAt * 1000).toISOString();
            return (
              <Grid.Column>
                <Link to={"/course:" + course.id}>
                  <Card>
                    <Card.Content>
                      <Card.Header>{course.title}</Card.Header>
                      <Card.Description>{course.summary}</Card.Description>
                      <Card.Meta>Created At {course.createdAt}</Card.Meta>
                      <Card.Meta>courseID : {course.id}</Card.Meta>
                    </Card.Content>
                  </Card>
                </Link>
              </Grid.Column>
            );
          })}
        {!courses.loading && yourClasses.length === 0 && (
          <Grid.Column>
            <Card>
              <Card.Content>
                <Card.Header>You have no classes yet.</Card.Header>
                {/* <Card.Description>{course.summary}</Card.Description>
               <Card.Meta>Created At {course.createdAt}</Card.Meta>
              <Card.Meta>courseID : {course.id}</Card.Meta> */}
              </Card.Content>
            </Card>
          </Grid.Column>
        )}
      </Grid>
      <Divider horizontal>
        <Header textAlign="left">
          <Icon name="video play" />
          Other Courses
        </Header>
      </Divider>
      <Grid columns={2} stackable>
        {!courses.loading &&
          otherClasses.map(course => {
            // let date = new Date(course.createdAt * 1000).toISOString();

            return (
              <Grid.Column>
                <Link to={"/course:" + course.id}>
                  <Card>
                    <Card.Content>
                      <Card.Header>{course.title}</Card.Header>
                      <Card.Description>{course.summary}</Card.Description>
                      <Card.Meta>Created At {course.createdAt}</Card.Meta>
                      <Card.Meta>courseID : {course.id}</Card.Meta>
                    </Card.Content>
                  </Card>
                </Link>
              </Grid.Column>
            );
          })}
        {!courses.loading && otherClasses.length === 0 && (
          <Grid.Column>
            <Card>
              <Card.Content>
                <Card.Header>You're not member of any classes.</Card.Header>
                {/* <Card.Description>{course.summary}</Card.Description>
               <Card.Meta>Created At {course.createdAt}</Card.Meta>
              <Card.Meta>courseID : {course.id}</Card.Meta> */}
              </Card.Content>
            </Card>
          </Grid.Column>
        )}
      </Grid>
    </Segment>
  );
}

export default Courses;
