import { gql, useQuery } from "@apollo/client";
import { faChalkboardTeacher } from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import { Link } from "react-router-dom";
import { Card, Grid, Segment } from "semantic-ui-react";
import Icon from "semantic-ui-react/dist/commonjs/elements/Icon";
import Image from "semantic-ui-react/dist/commonjs/elements/Image";
import Input from "semantic-ui-react/dist/commonjs/elements/Input";

const homePageBodyLStyle = {
    height: '100vh',
    backgroundColor: '#cadbeb',
    
}

const homePageCourseListLStyle = {
    backgroundColor: '#13678e',
    border: 'none',
    borderRadius: 7,
    margin: 'auto',
    height: '80%',
    overflow: 'auto',
    top: 80,
    margin: 40,
    marginTop: 10
}

const homepageSearchContainerStyle = {
    margin: 'auto',
    width: '50%',
    padding: '20px',
    top: 80,
    borderRadius: 5
}

const homepageSearchLStyle = {
    color: '#007fc1',
    top: 80,
    borderRadius: 5
    
}

const courseContentExtraLStyle = {
    color: '#023849',
    overFlow: 'hidden'
}

const courseDescriptionStyle = {
    overflow: 'hidden',
    height: '20px',
    borderRadius: 5
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
      
      tas {
        username
      }

      students {
        username
      }
    }
  }
`;


function Homepage() {
    const {data, loading, error} = useQuery(COURSES_QUERY, {
        variables: {
            keyWords: [],
            amount: 100,
            start: 0
        },
        fetchPolicy: "cache-and-network",
        nextFetchPolicy: "cache-first",
        onError(err) {
            console.log("error in getCourses:", err);
        }
    });

    console.log("data:", data);
    console.log("loading:", loading);
    console.log("error:", error);

    return (
        <div style={homePageBodyLStyle}>
            {/* <div style={homepageSearchContainerStyle}>
                <Input fluid icon='search' placeholder='Search for courses...' size='big' style={homepageSearchLStyle}/>
            </div> */}
            <Segment style={homePageCourseListLStyle}>
                <Grid columns={4}>
                    {!loading && data && 
                    data.coursesByKeyWords.map(course => {
                        let date = new Date(course.createdAt * 1000).toLocaleString("en-US", {
                            month: "long",
                            year: "numeric"
                        });
                        let imageSrc = 'https://source.unsplash.com/user/erondu'
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
                            <Grid.Column>
                                <Link to={"/course:" + course.id}>
                                    <Card
                                        onClick={() => {
                                        }}
                                        style={{width: '100vh'}}
                                    >
                                        {/*todo image handling*/}
                                        <Image src={imageSrc} wrapped ui={false}/>
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
                </Grid>
            </Segment>
        </div>
    );
}

export default Homepage;
