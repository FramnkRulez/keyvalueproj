import { render, screen } from '@testing-library/react';
import App from './App';

test('renders demo app header', () => {
  render(<App />);
  const linkElement = screen.getByText(/Demo App/i);
  expect(linkElement).toBeInTheDocument();
});
