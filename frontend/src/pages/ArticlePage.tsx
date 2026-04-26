import { RefreshCw } from 'lucide-react';
import { useCallback, useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Article, fetchArticle } from '../api';
import ArticleView from '../components/ArticleView';

export default function ArticlePage() {
  const { slug = '' } = useParams();
  const navigate = useNavigate();
  const [article, setArticle] = useState<Article | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  const load = useCallback(
    async (refresh = false) => {
      if (!slug) return;
      setLoading(true);
      setError(null);
      try {
        const result = await fetchArticle(slug, refresh);
        setArticle(result);
        if (result.slug !== slug) {
          navigate(`/wiki/${result.slug}`, { replace: true });
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Unable to generate article.');
      } finally {
        setLoading(false);
      }
    },
    [navigate, slug],
  );

  useEffect(() => {
    void load(false);
  }, [load]);

  if (loading) {
    return (
      <main className="article-shell">
        <div className="loading">
          <div className="spinner" />
          <p>Generating article...</p>
        </div>
      </main>
    );
  }

  if (error) {
    return (
      <main className="article-shell">
        <div className="error-state">
          <h1>Article generation failed</h1>
          <p>{error}</p>
          <button className="primary-button" onClick={() => void load(true)}>
            <RefreshCw size={16} aria-hidden="true" />
            Retry
          </button>
        </div>
      </main>
    );
  }

  if (!article) return null;

  return <ArticleView article={article} onRegenerate={() => void load(true)} />;
}
