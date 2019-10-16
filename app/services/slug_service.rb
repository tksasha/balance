# frozen_string_literal: true

class SlugService
  def initialize(word)
    @word = word
  end

  def build
    transliterate&.parameterize
  end

  private

  def transliterate
    TransliterateService.transliterate @word
  end

  class << self
    def build(word)
      new(word).build
    end
  end
end
