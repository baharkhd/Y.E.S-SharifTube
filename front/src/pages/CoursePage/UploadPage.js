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
import constants from "../../constants";
import FileUpload from "../FileUpload/FileUpload";
import _ from "lodash";

const fileUploadFrameLStyle = {
  borderColor: "#0021a3",
  margin: "auto",
  top: "120px",
  width: "70%",
  padding: "20px"
};

const OFFER_CONTENT_MUTATION = gql`
  mutation OfferContent(
    $courseID: String!
    $title: String!
    $description: String
    $video: Upload!
  ) {
    offerContent(
      courseID: $courseID
      target: {
        title: $title
        description: $description
        video: $video
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
        uploadedBY {
          username
        }
      }
      ... on Exception {
        message
      }
    }
  }
`;

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

const UPLOAD_MUTATION = gql`
  mutation UploadContent(
    $courseID: String!
    $title: String!
    $description: String
    $video: Upload!
    $tags: [String!]
  ) {
    uploadContent(
      courseID: $courseID
      target: {
        title: $title
        description: $description
        video: $video
        tags: $tags
      }
    ) {
      __typename
      ... on Content {
        id
        title
        description
        vurl
        approvedBY {
          name
          username
        }
        uploadedBY {
          name
          username
        }
        tags
        timestamp
      }
      ... on Exception {
        message
      }
    }
  }
`;

const UPLOAD_ATTACHMENTT_MUTATION = gql`
  mutation UploadAttachment(
    $courseID: String!
    $name: String!
    $attach: Upload!
    $description: String
  ) {
    uploadAttachment(
      courseID: $courseID
      target: { name: $name, attach: $attach, description: $description }
    ) {
      __typename
      ... on Attachment {
        id
        name
        description
        timestamp
        aurl
      }
      ... on Exception {
        message
      }
    }
  }
`;

function UploadPage(props) {
  let { courseID } = useParams();
  courseID = courseID.substring(1);

  const history = useHistory();

  let path = useLocation().pathname;
  let pathParts = path.split("/");
  var uploadType = pathParts[2];
  var fileType = pathParts[3];

  const [state, setState] = useState({
    title: "",
    description: "",
    url: "",
    tags: [],
    tagInput: "",
    file: ""
  });

  const [uploadAttachment] = useMutation(UPLOAD_ATTACHMENTT_MUTATION, {
    variables: {
      courseID: courseID,
      name: state.title,
      attach: state.file,
      description: state.description
    },
    update: (cache, { data: { uploadAttachment } }) => {
      const data = cache.readQuery({
        query: COURSE_QUERY,
        variables: {
          ids: [courseID]
        }
      });

      var localData = _.cloneDeep(data);

      localData.courses[0].inventory = localData.courses[0].inventory
        ? [...localData.courses[0].inventory, uploadAttachment]
        : [localData.courses[0].inventory];

      // console.log("localData:", localData);
      // console.log("uploadAttachment:", uploadAttachment);

      cache.writeQuery({
        query: COURSE_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ uploadAttachment }) => {
      console.log("upload attachmenttttttttt", uploadAttachment);
      if (uploadAttachment.__typename === "Attachment") {
        props.makeNotif(
          "Success!",
          "Attachment was successfully uploaded .",
          "success"
        );
      } else {
        props.makeNotif("Error!", offerContent.message, "danger");
      }
    }
  });

  const [uploadContent] = useMutation(UPLOAD_MUTATION, {
    variables: {
      courseID: courseID,
      title: state.title,
      description: state.description,
      video: state.file,
      tags: state.tags
    },
    update: (cache, { data: { uploadContent } }) => {
      const data = cache.readQuery({
        query: COURSE_QUERY,
        variables: {
          ids: [courseID]
        }
      });

      var localData = _.cloneDeep(data);
      localData.courses[0].contents = localData.courses[0].contents
        ? [...localData.courses[0].contents, uploadContent]
        : [uploadContent];

      cache.writeQuery({
        query: COURSE_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ uploadContent }) => {
      if (uploadContent.__typename === "Content") {
        props.makeNotif(
          "Success!",
          "Content was successfully uploaded .",
          "success"
        );
      } else {
        props.makeNotif("Error!", offerContent.message, "danger");
      }

      let path = "/course:" + courseID;
      history.push(path);
    }
  });

  const [offerContent] = useMutation(OFFER_CONTENT_MUTATION, {
    variables: {
      courseID: courseID,
      title: state.title,
      description: state.description,
      video: state.file
    },
    update: (cache, { data: { offerContent } }) => {
      const data = cache.readQuery({
        query: COURSE_QUERY,
        variables: {
          ids: [courseID]
        }
      });

      var localData = _.cloneDeep(data);

      localData.courses[0].pends = localData.courses[0].pends
        ? [...localData.courses[0].pends, offerContent]
        : [offerContent];

      cache.writeQuery({
        query: COURSE_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ offerContent }) => {
      console.log("*** offerContent:", offerContent);

      if (offerContent.__typename === "Pending") {
        props.makeNotif(
          "Success!",
          "Pending was successfully sent .",
          "success"
        );
      } else {
        props.makeNotif("Error!", offerContent.message, "danger");
      }

      let path = "/course:" + courseID;
      history.push(path);
    }
  });

  return (
    <Segment raised inverted style={fileUploadFrameLStyle}>
      {/* <Segment>Where you should upload videos</Segment> */}
      <Form inverted>
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
          style={{ resize: "none", height: "180px" }}
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
        {/* <FileUpload setFile={setState} otherState={state} /> */}
        <input
          type="file"
          onChange={e => {
            const [file] = e.target.files;

            console.log("-------------", file);
            setState({ ...state, file: file });
          }}
          style={{
            border: "none",
            paddingLeft: "0px",
            backgroundColor: "#1b1c1d",
            color: "white"
          }}
        />
        <Form.Button
          color="blue"
          onClick={() => {
            console.log("State before test:", state);

            if (
              state.title.trim() !== "" &&
              state.description.trim() !== "" &&
              state.file
            ) {
              if (uploadType == "upload") {
                if (fileType === "attachment") {
                  uploadAttachment();
                } else {
                  uploadContent();
                }
              } else {
                offerContent();
              }
            } else {
              props.makeNotif("Error!", constants.EMPTY_FIELDS, "danger");
            }
          }}
          style={{ marginTop: "10px" }}
        >
          Upload {fileType === "attachment" ? "Attachment" : "Video"}
        </Form.Button>
      </Form>
    </Segment>
  );
}

export default UploadPage;
