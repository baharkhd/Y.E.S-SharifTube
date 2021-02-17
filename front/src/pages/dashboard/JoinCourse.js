import React, {useState} from "react";
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
import {useMutation, gql, useQuery} from "@apollo/client";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChalkboardTeacher} from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import Icon from "semantic-ui-react/dist/commonjs/elements/Icon";

const joinCourseContainerHeaderStyle = {
    backgroundColor: "#4b4b4b"
}

const joinCourseSearchLStyle = {
    color: '#007fc1',
}

const courseContentExtraLStyle = {
    color: '#023849',
    overFlow: 'hidden'
}

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

const GET_COURSES_QUERY = gql`
  query GetCourses($ids: [String!]!) {
    courses(ids: $ids) {
      id
      title
      summary
      prof {
        username
        name
        emai
      }
      tas {
        username
        name
        email
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
        students {
          name
          username
        }
      }
      ... on Exception {
        message
      }
    }
  }
`;

function JoinCourseModel({joiningCourse, setState, username}) {
    const [courseInfo, setCourseInfo] = useState({
        courseID: "",
        token: ""
    });

    const {data, loading, error} = useQuery(COURSES_QUERY, {
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

        update(cache, {data: {addUserToCourse}}) {
            const data = cache.readQuery({
                query: GET_COURSES_QUERY,
                variables: {
                    ids: [courseInfo.courseID]
                }
            });

            console.log("addUserToCourse:", addUserToCourse);
            console.log("data in adding user to course:", data);
        },
        onCompleted: ({addUserToCourse}) => {
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
            <Modal.Header style={joinCourseContainerHeaderStyle}>
                <div>
                    <Input label="Join other classes" fluid autoFocus icon='search' placeholder='Search for courses...'
                           style={joinCourseSearchLStyle}/>
                </div>
            </Modal.Header>
            <Modal.Content scrolling>
                <Grid columns={3} stackable>
                    {!loading &&
                    data.coursesByKeyWords.map((course, i) => {
                        let date = new Date(course.createdAt * 1000).toLocaleString("en-US", {
                            month: "long",
                            year: "numeric",
                            day: "numeric"
                        });
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
                                    <Card.Content style={{overflow: 'hidden'}}>
                                        <Card.Header style={{overflow: 'hidden'}}>{course.title}</Card.Header>
                                        <Card.Meta>
                                            <span className='date'>{date}</span>
                                        </Card.Meta>
                                    </Card.Content>
                                    <Card.Content extra style={courseContentExtraLStyle}>
                                        <FontAwesomeIcon
                                            icon={faChalkboardTeacher}/>
                                        <span>&nbsp;&nbsp;</span>
                                        {/*todo click to see person account*/}
                                        {course.prof.username}
                                    </Card.Content>
                                </Card>
                                {courseInfo.courseID === course.id && (
                                    <Input
                                        autoFocus
                                        fluid
                                        placeholder={"password for class " + course.title}
                                        onChange={e => {
                                            setCourseInfo({...courseInfo, token: e.target.value});
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
                            // console.log("courseInfo.couresID:", courseInfo.courseID)
                            addUserToCourse();
                        }
                        setState({joiningCourse: false});
                    }}
                >
                    Join
                </Button>
                <Button
                    negative
                    onClick={() => {
                        setState({joiningCourse: false});
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
