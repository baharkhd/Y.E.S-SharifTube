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

// const UPLOAD_CONTENT_MUTATION = gql`
//   mutation UploadContent(
//     $courseID: String!
//     $title: String!
//     $description: String
//     $vurl: String!
//     $tags: [String!]
//   ) {
//     uploadContent(
//       courseID: $courseID
//       target: {
//         title: $title
//         description: $description
//         vurl: $vurl
//         tags: $tags
//       }
//     ) {
//       ... on Content {
//         id
//         title
//         description
//         vurl
//         tags
//         timestamp
//       }
//       ... on Exception {
//         message
//       }
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
        uploadedBY {
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

function UploadPage(props) {
  let { courseID } = useParams();
  courseID = courseID.substring(1);

  const history = useHistory();

  let path = useLocation().pathname;
  var n = path.lastIndexOf("/");
  var uploadType = path.substring(n + 1);

  const [state, setState] = useState({
    title: "",
    description: "",
    url: "",
    tags: [],
    tagInput: "",
    file: ""
  });

  // const [file, setFile] = useState(null)

  const [uploadContent] = useMutation(UPLOAD_MUTATION, {
    variables: {
      courseID: courseID,
      title: state.title,
      description: state.description,
      // video: {
      //   File: state.file,
      //   Filename: state.file.name,
      //   Size: state.file.size,
      //   ContentType: state.file.type
      // },
      video: state.file,
      tags: state.tags
    },
    onCompleted: ({ uploadContent }) => {
      console.log("updateContenttttttttt:", uploadContent);
    }
  });

  // console.log("+_+_+_+_+_+_+_+_+_+_+ uploadpage tokeeeeeeeeeeeeeeeeeeen:", localStorage.getItem(constants.AUTH_TOKEN))

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

  // const [uploadContent] = useMutation(UPLOAD_CONTENT_MUTATION, {
  //   variables: {
  //     courseID: courseID,
  //     title: state.title,
  //     description: state.description,
  //     vurl: state.url,
  //     tags: state.tags
  //   },
  //   update(cache, { data: { uploadContent } }) {
  //     var data = cache.readQuery({
  //       query: COURSE_QUERY,
  //       variables: {
  //         ids: [courseID]
  //       }
  //     });

  //     data = data[0];

  //     console.log("in update of uploadContent");
  //     console.log("uploadContent:", uploadContent);
  //     console.log("data:", data);
  //   },
  //   onCompleted: ({ uploadContent }) => {
  //     console.log("*** uploadContent:", uploadContent);
  //     let path = "/course:" + courseID;
  //     history.push(path);
  //   }
  // });

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
            console.log("State before test:", state)
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
      {/* <FileUpload setFile={setState} otherState={state} /> */}
      <input
        type="file"
        onChange={e => {
          const [file] = e.target.files
          
          console.log("-------------", file)
          setState({...state, file: file})
        }}
      />
    </Segment>
  );
}

export default UploadPage;
