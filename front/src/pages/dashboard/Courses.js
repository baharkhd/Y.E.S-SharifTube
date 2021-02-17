import React, {useState} from "react";
import {
    Grid,
    Segment,
    Image,
    Placeholder,
    Card,
    Divider,
    Header,
    Icon, Container
} from "semantic-ui-react";
import {Link} from "react-router-dom";
import {gql, useQuery} from "@apollo/client";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChalkboardTeacher} from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import {faChalkboard} from "@fortawesome/free-solid-svg-icons/faChalkboard";

const courseContentExtraLStyle = {
    color: '#023849',
    overFlow: 'hidden'
}

const courseDescriptionStyle = {
    overflow: 'hidden',
    height: '20px'
}

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
      createdAt
      
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
      
      students{
        username
      }
    }
  }
`;

function Courses(props) {
    const [courseIDs, setCourseIDs] = useState([]);

    const courses = useQuery(GET_COURSES_QUERY, {
        fetchPolicy: "cache-and-network",
        nextFetchPolicy: "cache-first",
        variables: {
            ids: courseIDs
        }
    });

    const {data, loading, error} = useQuery(GET_USER_QUERY, {
        fetchPolicy: "cache-and-network",
        nextFetchPolicy: "cache-first",
        onCompleted: ({user}) => {
            console.log("user:", user);
            setCourseIDs(user.courseIDs);
        }
    });

    console.log("coursesObject:", courses);
    console.log("coursesIDs", courseIDs);
    console.log("user courses:", data);
    console.log("loading:", loading);
    console.log("errror:", error);

    let yourClasses = [], otherClasses = [];

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
                     padding: 10,
                     height: '88vh',
                     overflowY: 'auto',
                     overflowX: 'hidden',
                     backgroundColor: "#ffffff",
                     borderColor: 'blue',
                     color: 'blue',
                 }}
        >
            <Divider horizontal>
                <Header textAlign="left">
                    <FontAwesomeIcon icon={faChalkboardTeacher} size='2x'/><span
                    style={{fontSize: '35px'}}>&nbsp;&nbsp;&nbsp;&nbsp;Your Courses</span>
                </Header>
            </Divider>
            <Grid columns={4} stackable style={{padding:'20px'}}>
                {!courses.loading &&
                yourClasses.map(course => {
                    let date = new Date(course.createdAt * 1000).toLocaleString("en-US", {
                        month: "long",
                        year: "numeric",
                        day: "numeric",
                        hour: "numeric",
                        minute: "numeric",
                    });
                    let memCount = 1
                    if (course.tas != null) {
                        memCount += course.tas.length
                    }
                    if (course.students != null) {
                        memCount += course.students.length
                    }
                    let memCountLabel = 'Members'
                    if (memCount === 1) {
                        memCountLabel = 'Member'
                    }
                    return (
                        <Grid.Column key={course.id}>
                            <Link to={"/course:" + course.id}>
                                <Card raised
                                    onClick={() => {
                                    }}
                                    style={{width: '100vh'}}>
                                    <Card.Content style={{overflow: 'hidden'}}>
                                        <Card.Header style={{overflow: 'hidden'}}>{course.title}</Card.Header>
                                        <Card.Meta>
                                            <span className='date'>{date}</span>
                                        </Card.Meta>
                                    </Card.Content>
                                    <Card.Content extra style={courseContentExtraLStyle}>
                                        <Icon name='user'/>
                                        {memCount} {memCountLabel}
                                    </Card.Content>
                                </Card>
                            </Link>
                        </Grid.Column>
                    );
                })}
                {!courses.loading && yourClasses.length === 0 && (
                    <Container textAlign="center" style={{marginTop: '40px'}}>
                        <Header as='h2' style={{color: 'red'}}>You have no classes yet.</Header>
                    </Container>
                )}
            </Grid>
            <Divider horizontal>
                <Header textAlign="left">
                    <FontAwesomeIcon icon={faChalkboard} size='2x'/><span
                    style={{fontSize: '35px'}}>&nbsp;&nbsp;&nbsp;&nbsp;Joined Courses</span>
                </Header>
            </Divider>
            <Grid columns={4} stackable style={{padding:'20px'}}>
                {!courses.loading &&
                otherClasses.map(course => {
                    let date = new Date(course.createdAt * 1000).toLocaleString("en-US", {
                        month: "long",
                        year: "numeric",
                    });
                    let memCount = 1
                    if (course.tas != null) {
                        memCount += course.tas.length
                    }
                    if (course.students != null) {
                        memCount += course.students.length
                    }
                    let memCountLabel = 'Members'
                    if (memCount === 1) {
                        memCountLabel = 'Member'
                    }
                    return (
                        <Grid.Column key={course.id}>
                            <Link to={"/course:" + course.id}>
                                <Card raised
                                      onClick={() => {
                                      }}
                                      style={{width: '100vh'}}>
                                    <Card.Content style={{overflow: 'hidden'}}>
                                        <Card.Header style={{overflow: 'hidden'}}>{course.title}</Card.Header>
                                        <Card.Meta>
                                            <span className='date'>{date}</span>
                                        </Card.Meta>
                                        <Card.Description
                                            style={courseDescriptionStyle}>{course.summary}</Card.Description>
                                    </Card.Content>
                                    <Card.Content extra style={courseContentExtraLStyle}>
                                        <FontAwesomeIcon
                                            icon={faChalkboardTeacher}/>
                                        <span>&nbsp;&nbsp;</span>
                                        {/*todo click to see person account*/}
                                        {course.prof.username}
                                        <br/>
                                        <Icon name='user'/>
                                        {memCount} {memCountLabel}
                                    </Card.Content>
                                </Card>
                            </Link>
                        </Grid.Column>
                    );
                })}
                {!courses.loading && otherClasses.length === 0 && (
                    <Container textAlign="center" style={{marginTop: '40px'}}>
                        <Header as='h2' style={{color: 'red'}}>You're not member of any class.</Header>
                    </Container>
                )}
            </Grid>
        </Segment>
    );
}

export default Courses;
