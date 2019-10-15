# frozen_string_literal: true

class TransliterateService
  include ActiveSupport::Inflector

  def initialize(word)
    @word = word
  end

  def transliterate
    return unless word.present?

    super word, ''
  end

  private

  def word
    return unless @word.present?

    @word
      .gsub(/Зг/, 'Zgh')
      .gsub(/зг/, 'zgh')
  end

  class << self
    def transliterate(word)
      new(word).transliterate
    end
  end
end
