import React from "react";
import {Card, Grid, Segment} from "semantic-ui-react";
import {Link} from "react-router-dom";
import {useMutation, gql, useQuery} from "@apollo/client";
import Image from "semantic-ui-react/dist/commonjs/elements/Image";
import Icon from "semantic-ui-react/dist/commonjs/elements/Icon";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChalkboardTeacher} from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import Input from "semantic-ui-react/dist/commonjs/elements/Input";
// import './Homepage.css'
const homePageBodyLStyle =
    {
        height: '100vh',
        backgroundColor: '#abe0fd'
    }

const homePageCourseListLStyle =
    {
        top: 70,
        backgroundColor: '#13678e',
        border: 'none',
        borderRadius: '0px'
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
            <Input icon='search' placeholder='Search...' className={Homepage.HomepageSearch}/>
            <Segment style={homePageCourseListLStyle}>
                <Grid columns={3}>
                    {!loading &&
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

                        return (
                            <Grid.Column>
                                <Link to={"/course:" + course.id}>
                                    <Card
                                        onClick={() => {
                                            console.log("course:", course);
                                        }}
                                    >
                                        {/*todo image handling*/}
                                        <Image src={imageSrc} wrapped ui={false}/>
                                        <Card.Content>
                                            <Card.Header>{course.title}</Card.Header>
                                            <Card.Meta>
                                                <span className='date'>{date}</span>
                                            </Card.Meta>
                                            <Card.Description>{course.summary}</Card.Description>
                                        </Card.Content>
                                        <Card.Content extra>
                                            <FontAwesomeIcon
                                                icon={faChalkboardTeacher}/><span>&nbsp;&nbsp;</span>{course.prof.username}
                                            <br/>
                                            <Icon name='user'/>
                                            {memCount} Members
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
