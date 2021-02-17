import React, { useState } from "react";
import {
  Segment,
  Label,
  Card,
  Icon,
  Grid,
  Button,
  Modal,
  Form,
  Input,
  TextArea,
  Message
} from "semantic-ui-react";
import { useParams } from "react-router-dom";
import { useQuery, gql, useMutation } from "@apollo/client";
import constants from "../../constants";

const PENDING_QUERY = gql`
  query GetPending(
    $courseID: String
    $status: Status
    $start: Int!
    $amount: Int!
  ) {
    pendings(
      filter: { courseID: $courseID, status: $status }
      start: $start
      amount: $amount
    ) {
      id
      title
      description
      furl
      status
      timestamp
      courseID
      uploadedBY {
        username
        name
      }
    }
  }
`;

const ACCEPT_PENDING_MUTATION = gql`
  mutation AcceptPending(
    $courseID: String!
    $pendingID: String!
    $title: String
    $description: String
    $tags: [String!]
    $message: String
  ) {
    acceptOfferedContent(
      courseID: $courseID
      pendingID: $pendingID
      changed: {
        title: $title
        description: $description
        tags: $tags
        message: $message
      }
    ) {
      ... on Pending {
        id
        title
        description
        furl
        status
        timestamp
      }
      ... on Exception {
        message
      }
    }
  }
`;

const REJECT_PENDING_MUTATION = gql`
  mutation RejectPending(
    $courseID: String!
    $pendingID: String!
    $message: String
  ) {
    rejectOfferedContent(
      courseID: $courseID
      pendingID: $pendingID
      message: { message: $message }
    ) {
      ... on Pending {
        id
        title
        description
        furl
        status
        timestamp
      }
      ... on Exception {
        message
      }
    }
  }
`;

const ChangePendingModal = props => {
  const [state, setState] = useState({
    title: props.title,
    description: props.description,
    tagInput: "",
    tags: []
  });

  const [acceptOfferedContent] = useMutation(ACCEPT_PENDING_MUTATION, {
    variables: {
      courseID: props.courseID,
      pendingID: props.pendingID,
      title: state.title,
      description: state.description,
      tags: state.tags,
      message: "test message for accept pending"
    },
    onCompleted: ({ acceptOfferedContent }) => {
      console.log("accept offered contenttttttt:", acceptOfferedContent);
      console.log("reject offered contentttttttt:", acceptOfferedContent);
      if (acceptOfferedContent.__typename === "Pending") {
        props.makeNotif("Success!", "Content was accepted!", "success");
      } else {
        props.makeNotif("Error!", acceptOfferedContent.message, "danger");
      }
    }
  });

  return (
    <Modal open={props.open}>
      <Modal.Header>Change the offered content if you want .</Modal.Header>

      <Modal.Content scrolling>
        <Modal.Content>
          <video width="60%" controls>
            <source
              src={
                // "https://s70.upera.net/2751313-0-WonderWoman4849193-480.mp4?owner=2640789&ref=1794068"
                props.furl
              }
              type="video/mp4"
            />
            {/* <source  type="" /> */}
            Your browser does not support HTML video.
          </video>
        </Modal.Content>
        <Form>
          <Form.Group>
            <Form.Field
              control={Input}
              label="Title"
              placeholder="Enter title for this content"
              value={state.title}
              onChange={e => {
                setState({ ...state, title: e.target.value });
              }}
            />
          </Form.Group>
          <Form.TextArea
            label="Description"
            placeholder="Enter description for this content"
            value={state.description}
            onChange={e => {
              setState({ ...state, description: e.target.value });
            }}
          />
          <Form.Group>
            <Form.Field
              control={Input}
              value={state.tagInput}
              // label="Tags"
              placeholder="Add a tag"
              onChange={e => {
                setState({ ...state, tagInput: e.target.value });
              }}
            />
            <Form.Field>
              <Form.Button
                icon="plus"
                positive
                onClick={() => {
                  if (state.tagInput !== "") {
                    setState({
                      ...state,
                      tags: [...state.tags, state.tagInput],
                      tagInput: ""
                    });
                  }
                }}
              />
            </Form.Field>
          </Form.Group>

          <Form.Field>
            <Label.Group>
              {state.tags.map(tag => {
                return (
                  <Label size="large">
                    <Icon name="hashtag" /> {tag}
                  </Label>
                );
              })}
            </Label.Group>
          </Form.Field>
        </Form>
      </Modal.Content>
      <Modal.Actions>
        <Button
          positive
          onClick={() => {
            if (state.title.trim() !== "" && state.description.trime() !== "") {
              acceptOfferedContent();
              props.setOpen(false);
            } else {
              props.makeNotif("Error!", constants.EMPTY_FIELDS, "danger");
            }
          }}
        >
          Change and Approve!
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
  tags,
  id,
  courseID,
  description,
  furl,
  makeNotif
}) => {
  let date = new Date(time * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });

  const [state, setState] = useState({
    modalOpen: false
    // contentID
  });

  const setOpen = val => {
    setState({ modalOpen: val });
  };

  const [rejectOfferedContent] = useMutation(REJECT_PENDING_MUTATION, {
    variables: {
      courseID: courseID,
      pendingID: id,
      message: "test message for reject pending"
    },
    onCompleted: ({ rejectOfferedContent }) => {
      console.log("reject offered contentttttttt:", rejectOfferedContent);
      if (rejectOfferedContent.__typename === "Pending") {
        makeNotif("Success!", "Content was rejected!", "success");
      } else {
        makeNotif("Error!", rejectOfferedContent.message, "danger");
      }
    }
  });

  return (
    <div>
      <ChangePendingModal
        pendingID={id}
        courseID={courseID}
        open={state.modalOpen}
        title={title}
        description={description}
        setOpen={setOpen}
        makeNotif={makeNotif}
        furl={furl}
      />
      <Card fluid>
        {/* <Card.Content>
          <Card.Header>{furl}</Card.Header>
        </Card.Content> */}
        <Card.Content>
          <Card.Header>{title}</Card.Header>
        </Card.Content>
        <Card.Content description>{description}</Card.Content>
        <Card.Content extra>
          uploaded by <b>{uploadedBY.name}</b> in <b>{date}</b>
        </Card.Content>

        <Card.Content>
          <Button.Group fluid>
            <Button
              positive
              onClick={() => {
                setState({ modalOpen: true });
              }}
            >
              Approve
            </Button>
            <Button
              color="red"
              onClick={() => {
                rejectOfferedContent();
              }}
            >
              Reject
            </Button>
          </Button.Group>
        </Card.Content>
      </Card>
    </div>
  );
};

function PendingPage(props) {
  let { courseID } = useParams();
  courseID = courseID.substring(1);

  // ($coureID: String, $status: Status, $uploaderUsername: String, $start: Int!, amount: Int!)

  const { data, loading, error } = useQuery(PENDING_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first",
    variables: {
      courseID: courseID,
      status: "PENDING",
      start: 0,
      amount: 100
    }
  });

  console.log("In pending page:");
  console.log("data:", data);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <Segment style={{ top: 70 }}>
      <Grid columns={3} stackable textAlign="left">
        {!loading &&
          (data.pendings != null && data.pendings.length !== 0 ? (
            data.pendings.map(content => {
              return (
                <Grid.Column textAlign="left">
                  <ContentCard
                    makeNotif={props.makeNotif}
                    title={content.title}
                    time={content.timestamp}
                    uploadedBY={content.uploadedBY}
                    description={content.description}
                    furl={content.furl}
                    // approvedBY={content.approvedBY}
                    // tags={content.tags}
                    id={content.id}
                    courseID={courseID}
                  />
                </Grid.Column>
              );
            })
          ) : (
            <Message warning size="massive" compact>
              <Message.Header>
                <Icon name="th" />
                There are no pendings yet
              </Message.Header>
            </Message>
          ))}
      </Grid>
    </Segment>
  );
}

export default PendingPage;
