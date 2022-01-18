import { PrismaClient } from ".prisma/client";

const client = new PrismaClient();

const organizationID = "organizationID";

client.project
  .create({
    data: {
      name: "Project 1",
      estimatedDuration: 24 * 30 * 3,
      startAt: new Date(),
      organization: {
        connect: {
          id: organizationID,
        },
      },
    },
  })
  .then((project) => {
    console.log(`project created id: ${project.id}`);
  })
  .catch((error) => console.error(error));
