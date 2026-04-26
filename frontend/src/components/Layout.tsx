import { Moon, Shuffle, Sun } from 'lucide-react';
import { ReactNode, useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { fetchRandomArticle } from '../api';
import SearchBar from './SearchBar';

type LayoutProps = {
  children: ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  const [dark, setDark] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    document.documentElement.dataset.theme = dark ? 'dark' : 'light';
  }, [dark]);

  async function openRandom() {
    const article = await fetchRandomArticle();
    navigate(`/wiki/${article.slug}`);
  }

  return (
    <div className="app-frame">
      <aside className="sidebar">
        <Link className="brand" to="/" aria-label="AIWIKI main page">
          <span className="brand-mark">A</span>
          <span>AIWIKI</span>
        </Link>
        <nav aria-label="Site">
          <Link to="/">Main page</Link>
          <button type="button" onClick={() => void openRandom()}>
            <Shuffle size={15} aria-hidden="true" />
            Random article
          </button>
          <Link to="/about">About AIWIKI</Link>
        </nav>
      </aside>

      <div className="main-column">
        <header className="topbar">
          <SearchBar />
          <button className="icon-button" onClick={() => setDark((value) => !value)} title="Toggle dark mode">
            {dark ? <Sun size={18} aria-hidden="true" /> : <Moon size={18} aria-hidden="true" />}
          </button>
        </header>
        {children}
        <footer>
          AIWIKI is an independent AI-generated demo. Not affiliated with Wikipedia, Wikimedia
          Foundation, MediaWiki, or Ollama. Verify important claims.
        </footer>
      </div>
    </div>
  );
}
