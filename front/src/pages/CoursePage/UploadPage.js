import React, { useState } from "react";
import {
  Segment,
  Input,
  Form,
  TextArea,
  Label,
  Button,
  Icon
} from "semantic-ui-react";
import { gql, useMutation } from "@apollo/client";
import { useParams, useHistory, useLocation } from "react-router-dom";

// uploadContent(username:String, courseID:String!, target:TargetContent!): UploadContentPayLoad!

// input TargetContent{
//   title: String!
//   description: String
//   vurl: String! # todo actual video
//   tags: [String!]
// }

// type Content{
//   id: ID!
//   title: String!
//   description: String
//   timestamp: Int!
//   uploadedBY: User!
//   approvedBY: User
//   vurl: String! #todo better implementation for video file
//   comments(start: Int!=0, amount: Int!=5): [Comment!]
//   tags: [String!]
//   courseID: String!
// }

const OFFER_CONTENT_MUTATION = gql`
  mutation OfferContent(
    $courseID: String!
    $title: String!
    $description: String
    $furl: String!
  ) {
    offerContent(
      courseID: $courseID
      target: {
        title: $title
        description: $description
        furl: $furl
        # tags: $tags
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

const UPLOAD_CONTENT_MUTATION = gql`
  mutation UploadContent(
    $courseID: String!
    $title: String!
    $description: String
    $vurl: String!
    $tags: [String!]
  ) {
    uploadContent(
      courseID: $courseID
      target: {
        title: $title
        description: $description
        vurl: $vurl
        tags: $tags
      }
    ) {
      ... on Content {
        id
        title
        description
        vurl
        tags
        timestamp
      }
      ... on Exception {
        message
      }
    }
  }
`;

function UploadPage(props) {
  const history = useHistory();

  let path = useLocation().pathname;
  var n = path.lastIndexOf("/");
  var uploadType = path.substring(n + 1);

  const [state, setState] = useState({
    title: "",
    description: "",
    url: "",
    tags: [],
    tagInput: ""
  });

  let { courseID } = useParams();
  courseID = courseID.substring(1);

  const [offerContent] = useMutation(OFFER_CONTENT_MUTATION, {
    variables: {
      courseID: courseID,
      title: state.title,
      description: state.description,
      furl: state.url
      // tags: state.tags
    },
    onCompleted: ({ offerContent }) => {
      console.log("*** offerContent:", offerContent);
      let path = "/course:" + courseID;
      history.push(path);
    }
  });

  const [uploadContent] = useMutation(UPLOAD_CONTENT_MUTATION, {
    variables: {
      courseID: courseID,
      title: state.title,
      description: state.description,
      vurl: state.url,
      tags: state.tags
    },
    onCompleted: ({ uploadContent }) => {
      console.log("*** uploadContent:", uploadContent);
      let path = "/course:" + courseID;
      history.push(path);
    }
  });

  return (
    <Segment style={{ top: 70 }}>
      {/* <Segment>Where you should upload videos</Segment> */}
      <Form>
        <Form.Group widths="four">
          <Form.Field
            control={Input}
            label="URL of this content"
            placeholder="URL"
            onChange={e => {
              setState({ ...state, url: e.target.value });
            }}
          />
        </Form.Group>
        <Form.Group widths="four">
          <Form.Field
            control={Input}
            label="Title of this content"
            placeholder="Title"
            onChange={e => {
              setState({ ...state, title: e.target.value });
            }}
          />
        </Form.Group>
        <Form.Field
          control={TextArea}
          label="Description of this content"
          placeholder="Write a summary about this content"
          onChange={e => {
            setState({ ...state, description: e.target.value });
          }}
        />

        {uploadType == "upload" ? (
          <div>
            <Form.Group>
              <Form.Field
                control={Input}
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
          </div>
        ) : (
          <></>
        )}
        <Form.Button
          color="blue"
          onClick={() => {
            if (uploadType == "upload") {
              uploadContent();
            } else {
              offerContent();
            }
          }}
        >
          Upload
        </Form.Button>
      </Form>
    </Segment>
  );
}

export default UploadPage;
