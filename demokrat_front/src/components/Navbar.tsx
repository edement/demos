
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link, useLocation } from 'react-router-dom';
import { Menu, X } from 'lucide-react';
import { cn } from '@/lib/utils';
import { useAuth } from '@/context/AuthContext';

const Navbar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const location = useLocation();
  const { user, logout, isAuthenticated } = useAuth();

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const closeMenu = () => {
    setIsMenuOpen(false);
  };

  const navigate = useNavigate();

  const isActive = (path: string) => {
    return location.pathname === path;
  };

  const navLinks = [
    { name: 'Главная', path: '/' },
    { name: 'Занятия', path: '/classes' },
    { name: isAuthenticated ? 'Профиль' : 'Войти', path: isAuthenticated ? '/dashboard' : '/auth' },
  ];

  return (
    <nav className="fixed top-0 left-0 w-full z-50 bg-demokrat-dark/95 border-b border-white/10 backdrop-blur-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          <div className="flex items-center">
            <Link to="/" onClick={closeMenu} className="flex-shrink-0">
              <span className="text-xl font-graffiti tracking-wider text-demokrat-purple">DEMOKRAT</span>
            </Link>
          </div>
          
          {/* Desktop Navigation */}
          <div className="hidden md:block">
            <div className="ml-10 flex items-center space-x-8">
              {navLinks.map((link) => (
                <Link
                  key={link.name}
                  to={link.path}
                  className={cn(
                    isActive(link.path)
                      ? 'text-demokrat-purple border-b-2 border-demokrat-purple'
                      : 'text-white hover:text-demokrat-purple',
                    'px-1 py-1 text-sm font-medium transition-colors'
                  )}
                >
                  {link.name}
                </Link>
              ))}
              {isAuthenticated && (
                <button
                  onClick={() => {
                    logout();
                    navigate('/');
                  }}
                  className="text-white hover:text-demokrat-purple px-1 py-1 text-sm font-medium transition-colors"
                >
                  Выйти
                </button>
              )}
            </div>
          </div>
          
          {/* Mobile Menu Button */}
          <div className="md:hidden">
            <button
              onClick={toggleMenu}
              className="inline-flex items-center justify-center p-2 rounded-md text-white hover:text-demokrat-purple focus:outline-none"
            >
              {isMenuOpen ? (
                <X className="block h-6 w-6" aria-hidden="true" />
              ) : (
                <Menu className="block h-6 w-6" aria-hidden="true" />
              )}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile Navigation */}
      {isMenuOpen && (
        <div className="md:hidden bg-demokrat-gray/95 backdrop-blur-sm border-b border-white/10">
          <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
            {navLinks.map((link) => (
              <Link
                key={link.name}
                to={link.path}
                onClick={closeMenu}
                className={cn(
                  isActive(link.path)
                    ? 'text-demokrat-purple'
                    : 'text-white hover:text-demokrat-purple',
                  'block px-3 py-2 text-base font-medium transition-colors'
                )}
              >
                {link.name}
              </Link>
            ))}
            {isAuthenticated && (
              <button
                onClick={() => {
                  logout();
                  closeMenu();
                  navigate('/');
                }}
                className="w-full text-left text-white hover:text-demokrat-purple block px-3 py-2 text-base font-medium transition-colors"
              >
                Выйти
              </button>
            )}
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;