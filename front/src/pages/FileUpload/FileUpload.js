import gql from "graphql-tag";
import React, { useCallback } from "react";
import { useDropzone } from "react-dropzone";
import { Icon, Segment } from "semantic-ui-react";


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
  const onDrop = useCallback(
    acceptedFiles => {
      const file = acceptedFiles[0];
      props.setFile({ ...props.otherState, file: file });
    }
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
