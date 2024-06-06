# IAM

Design an identity access management system for Different users, regions, progrms, API endpoints, and API methods.

## Requirements

### Functional Requirements

1. **User Roles and Access Control**:
   - Define distinct roles with specific permissions for admin, power user, editor, and general user.
   - Admins should have the capability to manage user accounts and change roles; power users should have extended privileges over general and editor functionalities, such as access to more sensitive system functions.
2. **Regional Compliance**:
   - Ensure data storage and processing complies with regional data residency laws, possibly requiring data to be stored and handled only in specific geographic locations.
3. **Program-Based Access Control**:
   - Implement access controls that restrict users based on the specific internal program or product they are authorized to use or manage.
4. **Authentication and Authorization**:
   - Multi-factor authentication (MFA) across all user types to enhance security.
   - Dynamic access control that can assess and respond to varying access needs in real-time, ensuring users only access what they are authorized to, based on their current role and the context of their access request.
5. **Audit and Reporting**:
   - Automated logging of all user activities, especially administrative actions, to support compliance with internal audits and regional regulations.
   - Detailed audit trails that can be used to review access and actions on sensitive data.
6. **API Security and Management**:
   - Secure API endpoints with OAuth 2.0 and potentially API gateways to manage throttling and secure access.
   - Define and enforce API usage policies tailored to different user roles and products.

### Non-Functional Requirements

1. **Scalability**:
   - The system must handle at least 10,000 users and 1,000 queries per hour, with the capability to scale as the number of internal users and transactions grows.
2. **Performance**:
   - Ensure that authentication and authorization processes are efficient, aiming for minimal latency to not adversely affect user experience despite the complex security mechanisms in place.
3. **High Availability**:
   - Design for high availability, particularly for authentication and authorization services, to minimize downtime and ensure consistent access for all users.
4. **Security**:
   - High-standard security practices including encryption of sensitive data at rest and in transit.
   - Regular security assessments to update and patch vulnerabilities, and ensure the integrity and confidentiality of user data.
5. **Compliance and Data Residency**:
   - The system must strictly comply with regional data protection laws and requirements, necessitating potentially multiple deployments in different regions.
6. **Maintainability and Extensibility**:
   - The system should be built with modern, maintainable code practices, allowing easy updates and feature additions.
   - Prepare for integration with other internal systems or potential expansions in user base or functionalities.

## High Level Design

### Components

1. **Authentication Service**

   - Manages user authentication processes including user login and session management.

   - Supports Multi-Factor Authentication (MFA) to provide an additional layer of security.

   - Use of JSON Web Tokens (JWT) to securely transmit information between parties as credentials.

2. **Authorization Service**

   - Dynamically adjusts access rights based on user attributes, context, and data sensitivity.

   - Support for Attribute-Based Access Control (ABAC) to allow more granular access control, complementing RBAC.

   - Caching frequently accessed policies to reduce decision latency and improve performance.

3. **User Management Service**

   - Provides APIs for creating, updating, and deleting user accounts.

   - Manages roles and permissions for each user.

   - Interface for administrative tasks such as user auditing, password resets, and role modifications.

4. **Audit Service**

   - Captures, stores, and indexes logs of all user activities and system actions.

   - Enhanced search and reporting capabilities for compliance monitoring and operational analytics.

   - Integration with external monitoring tools to trigger alerts on suspicious activities or breaches.

5. **API Gateway**

   - Acts as the single entry point for all incoming requests to the backend services.

   - Enforces security policies like rate limiting and request validation.

6. **Data Residency Service**

   - Manages data storage according to regional compliance requirements.

   - Ensures that data pertaining to a particular region is stored and processed within the same region.

   - Routing user data to the appropriate regional database.

   - Handling data replication and backup in compliance with local laws.

7. **Configuration and Discovery Service**

   - Manages service registration and discovery to ease the scaling of microservices.

   - Maintains environment-specific configurations that can be dynamically adjusted without service downtime.

   - Ensures all services can locate and communicate with each other in a secure and efficient manner.

### Workflow

1. **User Login:**

   - The user sends their credentials (username and password) or a previously issued refresh token to the Authentication Service via the `/login` endpoint.

   - The Authentication Service verifies the credentials against the user database. If multi-factor authentication (MFA) is required, this step includes that process.

   - Upon successful authentication, the service issues a JSON Web Token (JWT) that includes the user’s ID, roles, and any other necessary claims (e.g., region, program access).

2. **API Request Handling:**

   - The user, now authenticated, sends an API request that includes the JWT in the HTTP Authorization header.

   - The API Gateway first validates the JWT's signature to ensure it's still valid and checks its expiry. It applies any necessary rate limiting and performs initial request sanity checks (e.g., header validation, request size limits).

3. **Authorization:**

   - After initial validation, the API Gateway forwards the request to the designated microservice (e.g., User Management Service). Before processing, the microservice sends the JWT and request details (API endpoint and method) to the Authorization Service.

   - The Authorization Service validates the user’s roles and permissions against its permission mappings stored in a cache (for improved performance). These mappings detail permissible actions for each role across different endpoints, potentially filtered by region and program.

   - If the user lacks the necessary permissions, the Authorization Service responds with a "403 Forbidden." If permissions are validated, the flow proceeds.

4. **Fulfilling the Request:**

   - If the Authorization Service confirms the user has the appropriate permissions, the microservice processes the request (e.g., data retrieval, record updates).

   - Significant actions and data changes are logged asynchronously. Messages are sent to the Audit Service, which records them into the Audit Logs database.

5. **Response and Logging:**

   - The Audit Service logs all relevant information about the request, including the requester's identity, requested actions, and the timestamp.

   - The microservice completes processing the request and sends the response back to the user through the API Gateway.

## Deep Dive

### API Design

- Authentication Service: Endpoints in the Authentication Service manage user logins, sessions, and token operations

  - **POST `/auth/login`**: Authenticates a user's credentials and returns a JWT for successful logins.

  - **POST `/auth/refresh`**: Accepts a refresh token to issue a new JWT without requiring re-authentication.

  - **POST `/auth/logout`**: Invalidates the user's current session or token.

  - **POST `/auth/mfa`**: Handles multi-factor authentication challenges and verifications.

  - **GET `/auth/permissions/{userId}`**: Retrieves the permissions for a specific user ID.

  - **POST `/auth/verify`**: Checks if the provided token and requested action are permitted based on the user’s roles and permissions.

- User Management Service: This service manages user accounts, roles, and associated permissions:

  - **POST `/users`**: Creates a new user account.

  - **GET `/users/{userId}`**: Retrieves details for a specific user.

  - **PUT `/users/{userId}`**: Updates user account details.

  - **DELETE `/users/{userId}`**: Deletes a user account.

  - **POST `/users/{userId}/roles`**: Assigns roles to a user.

  - **DELETE `/users/{userId}/roles/{roleId}`**: Removes a role from a user.

  - **GET `/users/{userId}/roles`**: Lists all roles assigned to a user.

- Dynamic Role and Permission Management: These endpoints allow for real-time updates to roles and permissions:

  - **POST `/roles`**: Creates a new role with specific permissions.

  - **GET `/roles/{roleId}`**: Retrieves details for a specific role.

  - **PUT `/roles/{roleId}`**: Updates an existing role.

  - **DELETE `/roles/{roleId}`**: Deletes a role.

  - **POST `/permissions`**: Adds new permissions to the system.

  - **PUT `/permissions/{permissionId}`**: Modifies an existing permission.

  - **DELETE `/permissions/{permissionId}`**: Removes a permission from the system.

### Database Design

- **Table `users`**: Stores information about each user.
  - `user_id` (Primary Key, UUID): Unique identifier for the user.
  - `username` (VARCHAR): Unique username.
  - `email` (VARCHAR): User's email address.
  - `password_hash` (VARCHAR): Hashed password for security.
  - `created_at` (TIMESTAMP): The date and time when the account was created.
  - `updated_at` (TIMESTAMP): The date and time when the account was last updated.

- **Table `roles`**: Defines roles within the system.
  - `role_id` (Primary Key, UUID): Unique identifier for the role.
  - `role_name` (VARCHAR): Name of the role.
  - `description` (TEXT): Brief description of what the role entails.

- **Table `permissions`**: Defines specific actions that can be granted to roles.
  - `permission_id` (Primary Key, UUID): Unique identifier for the permission.
  - `permission_name` (VARCHAR): Name of the permission.
  - `description` (TEXT): Description of what the permission allows.

- **Table  `user_roles`**: Maps users to roles.
  - `user_id` (UUID, Foreign Key referencing `users.user_id`): Identifier for the user.
  - `role_id` (UUID, Foreign Key referencing `roles.role_id`): Identifier for the role.
  - Composite Primary Key (`user_id`, `role_id`).

- **Table `role_permissions`**: Maps roles to permissions.
  - `role_id` (UUID, Foreign Key referencing `roles.role_id`): Identifier for the role.
  - `permission_id` (UUID, Foreign Key referencing `permissions.permission_id`): Identifier for the permission.
  - Composite Primary Key (`role_id`, `permission_id`).

### User Management

```python
function handleRequest(user_id, api_endpoint, api_method, user_region, user_program):
    # Authenticate the user and get the user details
    user = authenticate(user_id)
    if user is None:
        log("Authentication failed for user_id: " + user_id)
        return "Authentication Failed", 401

    # Retrieve user roles based on the user's ID
    roles = getUserRoles(user_id)

    # Retrieve policies applicable to the user's roles, region, and program
    applicable_policies = getPolicies(roles, user_region, user_program)

    # Check if the request should be allowed based on the policies
    for policy in applicable_policies:
        # Check permissions for the given API endpoint and method
        if hasPermission(policy, api_endpoint, api_method):
            log("Access granted for user_id: " + user_id + " to " + api_endpoint + " using " + api_method)
            return "Access Granted", 200

    # If no policies allow the request, log and deny access
    log("Access denied for user_id: " + user_id + " to " + api_endpoint + " using " + api_method)
    return "Access Denied", 403
```

