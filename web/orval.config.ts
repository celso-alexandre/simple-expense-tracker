const openApiPath = `${__dirname}/../api/docs/swagger.yaml`;

export default {
  webApi: {
    input: openApiPath,
    output: {
      target: './src/util/api/generated/generated.ts',
      client: 'fetch',
      mode: 'split',
    },
    hooks: {
      afterAllFilesWrite: ['eslint --fix'],
    },
  },
};
