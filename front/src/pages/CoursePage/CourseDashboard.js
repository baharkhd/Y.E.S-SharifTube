import React, { useReducer, useCallback, useState } from "react";
import _ from "lodash";
import {
  Segment,
  Card,
  Sidebar,
  Label,
  Icon,
  Grid,
  Divider,
  Header,
  Search,
  Button
} from "semantic-ui-react";
import SideBar from "./CourseSidebar.js";
import "./CourseDashboard.css";
import { useParams, Link } from "react-router-dom";
import { gql, useQuery } from "@apollo/client";
import ContentsPart from "./ContentsPart";

// const GET_USER_QUERY = gql`
//   {
//     user {
//       username
//     }
//   }
// `;

const COURSE_QUERY = gql`
  query GetCoursesByID($ids: [String!]!) {
    courses(ids: $ids) {
      id
      title
      summary
      contents {
        id
        title
        description
        uploadedBY {
          name
          username
        }
        approvedBY {
          name
          username
        }
        tags
        timestamp
        vurl
      }
      prof {
        name
        username
        email
      }
      tas {
        name
        username
      }
      students {
        username
      }
      pends {
        title
        description
        id
      }
      inventory {
        id
        name
        aurl
        description
        timestamp
      }
    }
  }
`;

function CourseDashboard(props) {
  let { id } = useParams();
  id = id.substring(1);
  // Todo: use the course id to get the course information and use them

  console.log("---------------- username:", props.username);

  // const [contentsToSearch, setContents] = useState([]);

  const response = useQuery(COURSE_QUERY, {
    // fetchPolicy: "cache-and-network",
    // nextFetchPolicy: "cache-first",
    variables: {
      ids: [id]
    }
  });

  // const userResponse = useQuery(GET_USER_QUERY)

  let course = null;
  if (!response.loading && response.data) {
    course = response.data.courses[0];
    delete course.prof.__typename;
    if (course.tas) {
      course.tas.forEach(function(ta) {
        delete ta.__typename;
      });
    }
    // setContents(course.contents ? course.contents : [])
  }
  console.log("this course:", course);

  console.log(
    "data in course dashboard",
    response.data,
    response.loading,
    response.error
  );

  // let isProfTA;
  // if (!response.loading && course) {
  //   isProfTA =
  //     props.courseTAs.some(ta => ta.username === props.username) ||
  //     props.username === course.prof.username;
  // }

  return (
    <div>
      {!response.loading && course && (
        <SideBar
          isMobile={props.isMobile}
          sidebarIsOpen={props.sidebarOpen}
          // course={course}
          courseTitle={course.title}
          courseProf={course.prof}
          coursePends={course ? (course.pends ? course.pends : []) : []}
          courseTAs={course ? (course.tas ? course.tas : []) : []}
          username={props.username}
          isProf={props.username === course.prof.username} // can be prof or ta or st
          students={
            !response.loading
              ? response.data.courses[0].students
                ? response.data.courses[0].students
                : []
              : []
          }
        />
      )}
      <Segment
        style={{
          position: "absolute",
          left: props.isMobile ? 0 : 250,
          right: 0,
          margin: 30,
          top: 70
        }}
      >
        {!response.loading ? (
          course.contents ? (
            <ContentsPart contents={course.contents} id={id} />
          ) : (
            <div>
              <Divider horizontal>
                <Header textAlign="left">
                  <Icon name="video play" />
                  Videos
                </Header>
              </Divider>
              <Segment>There are no videos yet.</Segment>
            </div>
          )
        ) : (
          <></>
        )}
        <Divider horizontal>
          <Header textAlign="left">
            <Icon name="file" />
            Invetories
          </Header>
        </Divider>

        {!response.loading &&
        course &&
        course.inventory &&
        course.inventory.length !== 0 ? (
          <Grid columns={1} textAlign="left">
            {course.inventory.map(attach => {
              return (
                <Grid.Column>
                  <a
                    // href="https://sharif-webelopers.ir/static/images/background.jpg"
                    href={attach.aurl}
                    download
                  >
                    <Card>
                      <Card.Content>
                        <Card.Header>{attach.name}</Card.Header>
                        <Card.Description>
                          {attach.description}
                        </Card.Description>
                        <Card.Meta>aurl : {attach.aurl}</Card.Meta>
                      </Card.Content>
                    </Card>
                  </a>
                </Grid.Column>
              );
            })}
          </Grid>
        ) : (
          <Segment>There are no attachments yet .</Segment>
        )}
      </Segment>
    </div>
  );
}

export default CourseDashboard;
