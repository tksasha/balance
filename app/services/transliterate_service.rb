# frozen_string_literal: true

class TransliterateService
  include ActiveSupport::Inflector

  def initialize(word)
    @word = word
  end

  def transliterate
    return if word.blank?

    super word, ''
  end

  private

  def word
    return if @word.blank?

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
