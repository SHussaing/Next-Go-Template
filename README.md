# Next-Go-Template

This template repository provides a simple setup for running a **Next.js frontend** with a **Go backend**. The frontend is set up to automatically run the backend API using `concurrently`.


## Steps to Set Up and Run

1. **Navigate to the Frontend Directory**

   After creating the repository, go to the `frontend` directory:

   ```bash
   cd frontend
   ```

2. **Install Dependencies**

    Run the following command to install the required dependencies for the frontend:

   ```bash
   npm install
   ```
3. **Run the Application**

    To start the application, run:

   ```bash
   npm run dev
   ```

This will start both the Next.js frontend and the Go backend using the concurrently package, which ensures both the frontend and backend run together in the same terminal window.

The frontend will be available at http://localhost:3000.
The backend will be available at http://localhost:8080.


### Changing the Backend URL
If you need to change the backend URL (for example, when deploying or running on a different port), you will need to modify the .env.local file in the frontend directory. Update the NEXT_PUBLIC_GO_API variable to point to the new backend URL:

```bash
# change this to your backend api url
NEXT_PUBLIC_GO_API=http://new-backend-url:8080
```

### disclaimer
currently the .env is not being ignored in .gitignore