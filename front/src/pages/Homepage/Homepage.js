import React from "react";
import { Card, Grid, Segment } from "semantic-ui-react";
import { Link } from "react-router-dom";
import { useMutation, gql, useQuery } from "@apollo/client";

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

const COURSES_QUERY = gql`
  query GetCoursesByFilter($keyWords: [String!]!, $amount: Int!, $start: Int!) {
    coursesByKeyWords(keyWords: $keyWords, amount: $amount, start: $start) {
      id
      title
      summary
      createdAt
      # prof {
      #   username
      #   name
      # }
    }
  }
`;

function Homepage() {
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

  console.log("data:", data);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <Segment style={{ top: 70 }}>
      <Grid columns={3}>
        {!loading &&
          data.coursesByKeyWords.map(course => {
            let date = new Date(course.createdAt * 1000).toISOString();

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
                      <Card.Description>{course.summary}</Card.Description>
                      <Card.Meta>Created At {date}</Card.Meta>
                      <Card.Meta>courseID : {course.id}</Card.Meta>
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
