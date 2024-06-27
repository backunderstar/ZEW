import { type MockMethod } from 'vite-plugin-mock'
export default [
  {
    url: '/api/login',
    method: 'post',
    response: ({ body }: { body: { username: string; password: string } }) => {
      const { username, password } = body
      if (username === 'admin' && password === 'password123') {
        return {
          code: 200,
          message: 'Login successful',
          data: {
            token: 'your_generated_token_here',
            username: 'admin',
            role: 'admin'
          },
          success: true
        }
      } else {
        return {
          code: 401,
          message: 'Invalid username or password',
          success: false
        }
      }
    }
  },
  {
    url: '/api/menu',
    method: 'post',
    response: ({ token }: { token: string }) => {
      if (token === 'your_generated_token_here') {
        return {
          code: 200,
          message: 'Menu data retrieved successfully',
          data: [
            {
              id: 1,
              name: 'Dashboard',
              path: '/home/dashboard',
              icon: 'dashboard'
            },
            {
              id: 2,
              name: 'User',
              path: '/home/user',
              icon: 'User'
            },
            {
              id: 3,
              name: 'Article',
              path: '/home/article',
              icon: 'article'
            }
          ],
          success: true
        }
      } else {
        return {
          code: 401,
          message: 'Invalid token',
          success: false
        }
      }
    }
  }
] as MockMethod[]
