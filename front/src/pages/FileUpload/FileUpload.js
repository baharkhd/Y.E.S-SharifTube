import gql from "graphql-tag";
import React, { useCallback } from "react";
import { useDropzone } from "react-dropzone";
import { Icon, Segment } from "semantic-ui-react";

// uploadContent(username:String, courseID:String!, target:TargetContent!): UploadContentPayLoad!

// input TargetContent{
//     title: String!
//     description: String
//     video: Upload!
//     tags: [String!]
// }

// type Content{
//     id: ID!
//     title: String!
//     description: String
//     timestamp: Int!
//     uploadedBY: User!
//     approvedBY: User
//     vurl: String! #todo better implementation for video file
//     comments(start: Int!=0, amount: Int!=5): [Comment!]
//     tags: [String!]
//     courseID: String!
// }

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
        uploadedBY
      }
      ... on Exception {
        message
      }
    }
  }
`;

const FileUpload = props => {
  // const [file, setFile] = useState(null)

  //   const [uploadContent] = useMutation(UPLOAD_MUTATION);
  const onDrop = useCallback(
    acceptedFiles => {
      // select the first file from the Array of files
      const file = acceptedFiles[0];
      // use the uploadFile variable created earlier
      console.log("file:", file);
      //   setFile(file)
    //   console.log("my fileeeeee:", myFile)
      props.setFile({ ...props.otherState, file: file });
    }
    // ,
    // pass in uploadFile as a dependency
    // [uploadFile]
  );
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop
  });
  return (
    <>
      <div
        {...getRootProps()}
        className={`dropzone ${isDragActive && "isActive"}`}
      >
        <input {...getInputProps()} />
        {isDragActive ? (
          <Segment compact>Drop the files here ...</Segment>
        ) : (
          <Segment compact>
            <Icon name="upload" />
            Drag 'n' drop some files here, or click to select files
          </Segment>
        )}
      </div>
    </>
  );
};
export default FileUpload;
