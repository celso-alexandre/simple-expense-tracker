import 'dotenv/config';
import { join } from 'path';
import { exec } from 'child_process';

function generateSchemaSql() {
  const databaseUrl = process.env.DIRECT_URL?.split('?')[0];
  if (!databaseUrl) {
    console.error('DIRECT_URL environment variable is not defined.');
    return;
  }

  const rootDir = join(__dirname, '..');
  // console.log({ uri });
  const command = `
    docker run --rm --network="host" \
      -v ${rootDir}/prisma:/docker-entrypoint-initdb.d \
      postgres:16.3 \
      pg_dump "${databaseUrl}" -n public -s -f /docker-entrypoint-initdb.d/schema.sql
  `;

  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error generating schema.sql: ${error.message}`);
      return;
    }
    if (stderr) {
      console.error(`pg_dump error: ${stderr}`);
      return;
    }
    console.log('schema.sql file generated successfully!');
  });
}

generateSchemaSql();
