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

const pendingContents = [
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

// enum Status{
//   PENDING
//   ACCEPTED
//   REJECTED
// }

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

// pendings(filter: PendingFilter!, start: Int!=0, amount: Int!=5): [Pending!]!

// input PendingFilter{
//   courseID: String
//   status: Status
//   uploaderUsername: String
// }

const PENDING_QUERY = gql`
  query GetPending(
    $courseID: String
    $status: Status
    $uploaderUsername: String
    $start: Int!
    $amount: Int!
  ) {
    pendings(
      filter: {
        courseID: $courseID
        status: $status
        uploaderUsername: $uploaderUsername
      }
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
    <Modal open={true}>
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
  approvedBY,
  tags,
  id,
  courseID
}) => {
  return (
    <div>
      <Card fluid>
        <Card.Content>
          <Card.Header>{title}</Card.Header>
        </Card.Content>
        <Card.Content description>
          uploaded by <b>{uploadedBY}</b> and approved by <b>{approvedBY}</b> in
          time <b>{time}</b>
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
  });

  // ($coureID: String, $status: Status, $uploaderUsername: String, $start: Int!, amount: Int!)

  const { data, loading, error } = useQuery(PENDING_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first",
    variables: {
      courseID,
      status: "PENDING",
      uploaderUsername: props.username,
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
      <ChangePendingModal />
      <Grid columns={3} stackable>
        {pendingContents.map(content => {
          return (
            <Grid.Column textAlign="left">
              <ContentCard
                title={content.title}
                time={content.timestamp}
                uploadedBY={content.uploadedBY}
                approvedBY={content.approvedBY}
                tags={content.tags}
                id={content.id}
                courseID={courseID}
              />
            </Grid.Column>
          );
        })}
      </Grid>
    </Segment>
  );
}

export default PendingPage;
