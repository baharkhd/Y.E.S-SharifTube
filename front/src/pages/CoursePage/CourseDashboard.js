import React, { useReducer, useCallback } from "react";
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

const contents = [
  {
    title: "title1",
    timestamp: "time1",
    uploadedBY: "uploadedBy1",
    approvedBY: "approvedBy1",
    tags: ["tags1-1-sdfkjnsd", "tags1-2", "tags1-3"],
    id: "videoID1"
  },
  {
    title: "title2",
    timestamp: "time2",
    uploadedBY: "uploadedBy2",
    approvedBY: "approvedBy2",
    tags: ["tags2-1", "tags2-2", "tags2-3"],
    id: "videoID2"
  },
  {
    title: "title3",
    timestamp: "time3",
    uploadedBY: "uploadedBy3",
    approvedBY: "approvedBy3",
    tags: ["tags3-1", "tags3-2", "tags3-3"],
    id: "videoID3"
  },
  {
    title: "title4",
    timestamp: "time4",
    uploadedBY: "uploadedBy4",
    approvedBY: "approvedBy4",
    tags: ["tags4-1", "tags4-2", "tags4-3"],
    id: "videoID4"
  },
  {
    title: "title5",
    timestamp: "time5",
    uploadedBY: "uploadedBy5",
    approvedBY: "approvedBy5",
    tags: ["tags5-1", "tags5-2", "tags5-3"],
    id: "videoID5"
  },
  {
    title: "title6",
    timestamp: "time6",
    uploadedBY: "uploadedBy6",
    approvedBY: "approvedBy6",
    tags: ["tags6-1", "tags6-2", "tags6-3"],
    id: "videoID6"
  }
];

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
    }
  }
`;

const initialState = {
  loading: false,
  results: [],
  value: "",
  resultsShown: []
};

function searchReducer(state, action) {
  switch (action.type) {
    case "CLEAN_QUERY":
      return initialState;
    case "START_SEARCH":
      return { ...state, loading: true, value: action.query };
    case "FINISH_SEARCH":
      return { ...state, loading: false, results: action.results };
    case "UPDATE_SELECTION":
      return { ...state, value: action.selection };
    case "UPDATE_SHOWN_RESULTS":
      return { ...state, resultsShown: action.data };
    default:
      throw new Error();
  }
}

const ContentCard = ({
  title,
  time,
  uploadedBY,
  approvedBY,
  tags,
  id,
  courseID
}) => {
  return (
    <div>
      <Link to={"/course:" + courseID + "/content:" + id}>
        <Card fluid className="Content">
          <Card.Content>
            <Card.Header>{title}</Card.Header>
          </Card.Content>
          <Card.Content description>
            uploaded by <b>{uploadedBY}</b> and approved by <b>{approvedBY}</b>{" "}
            in time <b>{time}</b>
          </Card.Content>
          <Card.Content extra>
            {tags.map(tag => {
              return (
                <Label style={{ marginBottom: 5 }}>
                  <Icon name="hashtag" /> {tag}
                </Label>
              );
            })}
          </Card.Content>
        </Card>
      </Link>
    </div>
  );
};

function CourseDashboard(props) {
  let { id } = useParams();
  id = id.substring(1);
  // Todo: use the course id to get the course information and use them

  console.log("---------------- username:", props.username);

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
  }
  console.log("this course:", course);

  console.log(
    "data in course dashboard",
    response.data,
    response.loading,
    response.error
  );

  const [state, dispatch] = React.useReducer(searchReducer, initialState);
  const { loading, results, value, resultsShown } = state;

  if (resultsShown.length == 0) {
    dispatch({ type: "UPDATE_SHOWN_RESULTS", data: contents });
  }

  const timeoutRef = React.useRef();
  const handleSearchChange = React.useCallback((e, data) => {
    clearTimeout(timeoutRef.current);
    dispatch({ type: "START_SEARCH", query: data.value });

    timeoutRef.current = setTimeout(() => {
      if (data.value.length === 0) {
        dispatch({ type: "CLEAN_QUERY" });
        return;
      }

      const re = new RegExp(_.escapeRegExp(data.value), "i");
      // const isMatch = result => re.test(result.title);

      const isMatch = result => {
        var tag;
        console.log("tags:", result.tags);
        for (tag of result.tags) {
          let check = re.test(tag);
          console.log("checkkkk:", check, ", tag:", tag);
          if (check) {
            return true;
          }
        }
        return false;

        // result.tags.every(tag => {
        //   // console.log("tag:", tag)
        //   let check = re.test(tag)f
        //   // console.log("is match:", check)

        // });
        // return true;
      };

      let newArray = _.filter(contents, function(ss) {
        // console.log("ss:", ss);
        let check = isMatch(ss);
        // console.log("check:", check);
        return check;
      });
      // console.log("newArray:", newArray);

      dispatch({
        type: "FINISH_SEARCH",
        results: newArray
      });
    }, 300);
  }, []);
  React.useEffect(() => {
    return () => {
      clearTimeout(timeoutRef.current);
    };
  }, []);

  return (
    <div>
      {!response.loading && (
        <SideBar
          isMobile={props.isMobile}
          sidebarIsOpen={props.sidebarOpen}
          // course={course}
          courseTitle={course.title}
          courseProf={course.prof}
          courseTAs={course ? (course.tas ? course.tas : []) : []}
          username={props.username}
          role={props.username === course.prof.username ? "prof" : "non-prof"} // can be prof or ta or st
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
        <Search
          // aligned
          loading={loading}
          onResultSelect={(e, data) => {
            dispatch({
              type: "UPDATE_SELECTION",
              selection: data.result.title
            });
            // setValue(data.result.title);
          }}
          onSearchChange={handleSearchChange}
          results={results}
          value={value}
        />
        <Button
          onClick={() => {
            dispatch({ type: "UPDATE_SHOWN_RESULTS", data: results });
          }}
        >
          check
        </Button>
        <Divider horizontal>
          <Header textAlign="left">
            <Icon name="video play" />
            Videos
          </Header>
        </Divider>
        <Grid columns={2} stackable>
          {resultsShown.map(content => {
            return (
              <Grid.Column textAlign="left">
                <ContentCard
                  title={content.title}
                  time={content.timestamp}
                  uploadedBY={content.uploadedBY}
                  approvedBY={content.approvedBY}
                  tags={content.tags}
                  id={content.id}
                  courseID={id}
                />
              </Grid.Column>
            );
          })}
        </Grid>
        <Divider horizontal>
          <Header textAlign="left">
            <Icon name="file" />
            Invetories
          </Header>
        </Divider>
        <Grid columns={1} textAlign="left">
          <Grid.Column>Sample inventory</Grid.Column>
        </Grid>
      </Segment>
    </div>
  );
}

export default CourseDashboard;
