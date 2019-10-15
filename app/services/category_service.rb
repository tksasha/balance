# frozen_string_literal: true

class CategoryService
  delegate :errors, to: :category

  def initialize(params)
    @params = params
  end

  def save
    category.save
  end

  private

  def category
    @category ||= Category.new params
  end

  def slug
    TransliterateService.transliterate(@params[:name])&.parameterize
  end

  def params
    @params.merge(slug: slug)
  end
end
