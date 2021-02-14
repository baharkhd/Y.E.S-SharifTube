import React, { useState } from "react";
import { Grid, Segment, Image, Placeholder } from "semantic-ui-react";
import { gql, useQuery } from "@apollo/client";

const GET_USER_QUERY = gql`
  {
    user {
      username
      name
      password
      email
      courseIDs
    }
  }
`;

const GET_COURSES_QUERY = gql`
  query GetCourses($ids: [String!]!) {
    courses(ids: $ids) {
      id
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

  console.log("coursesIDs", courseIDs);
  console.log("user courses:", data);
  console.log("loading:", loading);
  console.log("errror:", error);

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
      <Grid columns={3} stackable>
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
