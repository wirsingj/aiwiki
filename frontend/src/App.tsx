import { Route, Routes } from 'react-router-dom';
import Layout from './components/Layout';
import ArticlePage from './pages/ArticlePage';
import Home from './pages/Home';

export default function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/wiki/:slug" element={<ArticlePage />} />
        <Route
          path="/about"
          element={
            <main className="article-shell">
              <h1>About AIWIKI</h1>
              <p>
                AIWIKI is an independent local demo encyclopedia that drafts articles with your
                configured model. It is inspired by classic wiki layouts, but it is not affiliated
                with Wikipedia, the Wikimedia Foundation, MediaWiki, Ollama, or any model provider,
                and it does not use scraped encyclopedia content.
              </p>
              <p className="notice">
                AI-generated articles may contain inaccuracies. Generated notes are not verified
                citations. Verify important claims and comply with the licenses for the models you run.
              </p>
            </main>
          }
        />
      </Routes>
    </Layout>
  );
}
