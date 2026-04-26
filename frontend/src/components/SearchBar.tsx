import { Search } from 'lucide-react';
import { FormEvent, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { safeTopicSlug, validateTopicInput } from '../api';

type SearchBarProps = {
  size?: 'normal' | 'large';
};

export default function SearchBar({ size = 'normal' }: SearchBarProps) {
  const [topic, setTopic] = useState('');
  const [error, setError] = useState<string | null>(null);
  const location = useLocation();
  const navigate = useNavigate();
  const currentTopic = size === 'normal' ? topicFromWikiPath(location.pathname) : '';

  function submit(event: FormEvent) {
    event.preventDefault();
    const validationError = validateTopicInput(topic);
    if (validationError) {
      setError(validationError);
      return;
    }

    const slug = safeTopicSlug(topic);
    navigate(`/wiki/${slug}`);
    setTopic('');
    setError(null);
  }

  return (
    <form className={`search search-${size}`} onSubmit={submit} noValidate>
      <label className="sr-only" htmlFor={`search-${size}`}>
        Search AIWIKI
      </label>
      <div className="search-row">
        <input
          id={`search-${size}`}
          value={topic}
          onChange={(event) => {
            setTopic(event.target.value);
            if (error) setError(null);
          }}
          placeholder={currentTopic || 'Search AIWIKI'}
          autoComplete="off"
          maxLength={121}
          aria-invalid={Boolean(error)}
          aria-describedby={error ? `search-${size}-error` : undefined}
        />
        <button type="submit" title="Search">
          <Search size={18} aria-hidden="true" />
          <span>Search</span>
        </button>
      </div>
      {error && (
        <p className="search-error" id={`search-${size}-error`}>
          {error}
        </p>
      )}
    </form>
  );
}

function topicFromWikiPath(pathname: string): string {
  const match = pathname.match(/^\/wiki\/([^/]+)/);
  if (!match) return '';

  try {
    return decodeURIComponent(match[1]).replace(/-/g, ' ');
  } catch {
    return match[1].replace(/-/g, ' ');
  }
}
