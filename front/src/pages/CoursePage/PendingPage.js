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
  TextArea
} from "semantic-ui-react";
import { useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";

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

// input EditedPending{
//   title: String
//   description: String
// }

// acceptOfferedContent(username:String, courseID:String!, pendingID:String!, changed:EditedPending!): EditOfferedContentPayLoad!

// rejectOfferedContent(username:String, courseID:String!, pendingID:String!): DeleteOfferedContentPayLoad!

const ACCEPT_PENDING_MUTATION = gql`
  mutation AcceptPending(
    $courseID: String!
    $pendingID: String!
    $title: String
    $description: String
  ) {
    acceptOfferedContent(
      courseID: $courseID
      pendingID: $pendingID
      changed: { title: $title, description: $description }
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
  mutation RejectPending($courseID: String!, $pendingID: String!) {
    rejectOfferedContent(courseID: $courseID, pendingID: $pendingID) {
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

// type Pending{
//   id: ID!
//   title: String!
//   description: String
//   status: Status!
//   timestamp:Int!
//   uploadedBY: User!
//   furl: String! #todo better implementation for file
//   courseID: String!
// }

const ChangePendingModal = props => {
  const [state, setState] = useState({
    title: props.title,
    description: props.description,
    tagInput: "",
    tags: []
  });
  return (
    <Modal open={props.open}>
      <Modal.Header>Change the content if you want :)</Modal.Header>
      <Modal.Content>
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
                      tags: [...state.tags, state.tagInput]
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
        <Button positive>Change and Approve!</Button>
        <Button negative>Cancel</Button>
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
  furl
}) => {
  let date = new Date(time * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });
  return (
    <div>
      <Card fluid>
        <Card.Content>
          <Card.Header>{furl}</Card.Header>
        </Card.Content>
        <Card.Content>
          <Card.Header>{title}</Card.Header>
        </Card.Content>
        <Card.Content description>{description}</Card.Content>
        <Card.Content extra>
          uploaded by <b>{uploadedBY.name}</b> in <b>{date}</b>
        </Card.Content>

        <Card.Content>
          <Button.Group fluid>
            <Button positive>Approve</Button>
            {/* <Button color="blue">salam</Button> */}
            <Button color="red">Reject</Button>
          </Button.Group>
        </Card.Content>
      </Card>
    </div>
  );
};

function PendingPage(props) {
  let { courseID } = useParams();
  courseID = courseID.substring(1);

  const [state, setState] = useState({
    modalOpen: false
    // contentID
  });

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
      <ChangePendingModal open={state.modalOpen} />
      <Grid columns={3} stackable>
        {!loading &&
          (data.pendings ? (
            data.pendings.map(content => {
              return (
                <Grid.Column textAlign="left">
                  <ContentCard
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
            <Segment>There are no pending contents yet</Segment>
          ))}
      </Grid>
    </Segment>
  );
}

export default PendingPage;
