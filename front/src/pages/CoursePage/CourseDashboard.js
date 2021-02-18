import { gql, useQuery } from "@apollo/client";
import React from "react";
import { useParams } from "react-router-dom";
import {Card, Container, Divider, Grid, Header, Icon, Segment} from "semantic-ui-react";
import ContentsPart from "./ContentsPart";
import "./CourseDashboard.css";
import SideBar from "./CourseSidebar.js";


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

  const response = useQuery(COURSE_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first",
    variables: {
      ids: [id]
    }
  });


  let course = null;
  if (!response.loading && response.data) {
    course = response.data.courses[0];
    delete course.prof.__typename;
    if (course.tas) {
      course.tas.forEach(function(ta) {
        delete ta.__typename;
      });
    }
  }

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
            top: 70,
            height: '85vh',
            overflowY: 'auto',
            padding: '20px',
            borderColor: '#042b61'
        }}
      >
        {!response.loading ? (
          course.contents ? (
            <ContentsPart contents={course.contents} id={id} makeNotif={props.makeNotif} />
          ) : (
            <div>
              <Divider horizontal>
                <Header textAlign="left">
                  <Icon name="video play" />
                  Videos
                </Header>
              </Divider>
                <Container textAlign="center" style={{marginTop: '40px'}}>
                    <Header as='h2' style={{color: 'red'}}>There are no videos yet.</Header>
                </Container>
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
          <Grid columns={2} textAlign="left">
            {course.inventory.map(attach => {
              return (
                <Grid.Column>
                  <a
                    href={attach.aurl}
                    download
                  >
                    <Card fluid className="Attachment">
                      <Card.Content>
                        <Card.Header>{attach.name}</Card.Header>
                        <Card.Description>
                          {attach.description}
                        </Card.Description>
                      </Card.Content>
                    </Card>
                  </a>
                </Grid.Column>
              );
            })}
          </Grid>
        ) : (
            <Container textAlign="center" style={{marginTop: '40px'}}>
                <Header as='h2' style={{color: 'red'}}>There are no attachments yet.</Header>
            </Container>
        )}
      </Segment>
    </div>
  );
}

export default CourseDashboard;
