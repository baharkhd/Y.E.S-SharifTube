import React, { useCallback } from "react";
import { useDropzone } from "react-dropzone";
// import { useMutation } from "@apollo/client";
import { useMutation } from "@apollo/client";
import gql from "graphql-tag";
import { Segment, Input, Icon } from "semantic-ui-react";

const UploadMutation = gql`
  mutation uploadFile($file: Upload!) {
    uploadFile(file: $file) {
      path
      id
      filename
      mimetype
    }
  }
`;
// pass in the UploadMutation mutation we created earlier.
const FileUpload = () => {
  const [uploadFile] = useMutation(UploadMutation);
  const onDrop = useCallback(
    acceptedFiles => {
      // select the first file from the Array of files
      const file = acceptedFiles[0];
      // use the uploadFile variable created earlier
      console.log("file:", file);
      //   uploadFile({
      //     // use the variables option so that you can pass in the file we got above
      //     variables: { file },
      //     onCompleted: () => {},
      //   });
    },
    // pass in uploadFile as a dependency
    [uploadFile]
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
          <Segment compact >
            Drop the files here ...
          </Segment>
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
