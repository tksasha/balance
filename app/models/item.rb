class Item < ActiveRecord::Base
  include ActsAsHasFormula

  belongs_to :category

  validates :date, :category_id, :formula, presence: true

  scope :income, -> { joins(:category).where({ categories: { income: true } }) }

  scope :expense, -> { joins(:category).where({ categories: { income: false } }) }

  acts_as_paranoid

  class << self
    #
    # TODO: rename it to `search_by` and refactor it
    # TODO: use ItemSearcher instead
    #
    def search date_range, slug=nil
      if slug
        where 'categories.slug' => slug
      else
        self
      end.includes(:category).where(date: date_range).order('date DESC')
    end
  end
end
