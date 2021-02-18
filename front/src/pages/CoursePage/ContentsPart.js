import { gql, useMutation } from "@apollo/client";
import _ from "lodash";
import React, { useState } from "react";
import { Link } from "react-router-dom";
import { Button, Card, Divider, Grid, Header, Icon, Label, Modal, Search, Segment } from "semantic-ui-react";

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

const DELETE_CONTENT = gql`
  mutation DeleteContent($courseID: String!, $contentID: String!) {
    deleteContent(courseID: $courseID, contentID: $contentID) {
      ... on Content {
        id
        title
        description
      }
      ... on Exception {
        message
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

const DeletModal = props => {
  const [deleteContent] = useMutation(DELETE_CONTENT, {
    update: (cache, { data: { deleteContent } }) => {
      const data = cache.readQuery({
        query: COURSE_QUERY,
        variables: {
          ids: [props.courseID]
        }
      });

      var localData = _.cloneDeep(data);
      localData.courses[0].contents = localData.courses[0].contents.filter(
        content => {
          return content.id !== deleteContent.id;
        }
      );

      cache.writeQuery({
        query: COURSE_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ deleteContent }) => {
      if (deleteContent.__typename === "Content") {
        props.makeNotif("Success", "Content successfully removed .", "success");
      } else {
        props.makeNotif("Error", deleteContent.message, "danger");
      }
    }
  });

  return (
    <Modal open={props.open}>
      <Modal.Header>Are you sure you want to delete this content?</Modal.Header>
      <Modal.Actions>
        <Button
          positive
          onClick={() => {
            deleteContent({
              variables: {
                courseID: props.courseID,
                contentID: props.contentID
              }
            });
            props.setOpen(false);
          }}
        >
          Delete
        </Button>
        <Button
          negative
          onClick={() => {
            props.setOpen(false);
          }}
        >
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
};

const ContentCard = ({
  title,
  time,
  uploadedBY,
  approvedBY,
  tags,
  id,
  courseID,
  isSearched,
  makeNotif
}) => {
  let date = new Date(time * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });

  const [open, setOpen] = useState(false);

  return (
    <div>
      <DeletModal
        courseID={courseID}
        contentID={id}
        makeNotif={makeNotif}
        open={open}
        setOpen={setOpen}
      />
      <Segment style={{ backgroundColor: "#d3dfed" }}>
        <Link to={"/course:" + courseID + "/content:" + id}>
          <Card
            fluid
            className="Content"
            style={{ backgroundColor: isSearched ? "#fffa63" : "" }}
          >
            <Card.Content>
              <Card.Header>{title}</Card.Header>
            </Card.Content>
            <Card.Content description>
              uploaded by <b>{uploadedBY.name}</b>
              {approvedBY ? " and approved by" : ""}{" "}
              <b>{approvedBY ? approvedBY.name : ""}</b> in time <b>{date}</b>
            </Card.Content>
            <Card.Content extra>
              {tags &&
                tags.map(tag => {
                  return (
                    <Label style={{ marginBottom: 5 }}>
                      <Icon name="hashtag" /> {tag}
                    </Label>
                  );
                })}
            </Card.Content>
          </Card>
        </Link>
        <Divider />
        <Button
          icon="x"
          color="black"
          onClick={() => {
            setOpen(true);
          }}
        />
      </Segment>
    </div>
  );
};

function ContentsPart({ contents, id, makeNotif }) {
  const [state, dispatch] = React.useReducer(searchReducer, initialState);
  const { loading, results, value, resultsShown } = state;

  if (resultsShown.length == 0) {
    dispatch({
      type: "UPDATE_SHOWN_RESULTS",
      data: contents
    });
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

      const isMatch = result => {
        var tag;
        if (result.tags) {
          for (tag of result.tags) {
            let check = re.test(tag);
            if (check) {
              return true;
            }
          }
          return false;
        } else {
          return false;
        }
      };

      let newArray = _.filter(contents, function(ss) {
        let check = isMatch(ss);
        return check;
      });

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
      <Grid columns={2} stackable>
        <Grid.Column floated="left">
          <Search
            showNoResults={false}
            loading={loading}
            onResultSelect={(e, data) => {
              dispatch({
                type: "UPDATE_SELECTION",
                selection: data.result.title
              });
            }}
            onSearchChange={handleSearchChange}
            results={[]}
            value={value}
            placeholder="Search for tags ..."
          />
        </Grid.Column>
      </Grid>

      <Divider horizontal>
        <Header textAlign="left">
          <Icon name="video play" />
          Videos
        </Header>
      </Divider>
      <Grid columns={2} stackable>
        {resultsShown.map(content => {
          let isSearched = results.includes(content, c => {
            return c.id === content.id;
          });
          console.log("isSearched:", isSearched);
          return (
            <Grid.Column textAlign="left">
              <ContentCard
                key={content.id}
                title={content.title}
                time={content.timestamp}
                uploadedBY={content.uploadedBY}
                approvedBY={content.approvedBY}
                tags={content.tags}
                id={content.id}
                courseID={id}
                isSearched={isSearched}
                makeNotif={makeNotif}
              />
            </Grid.Column>
          );
        })}
      </Grid>
    </div>
  );
}

export default ContentsPart;
