class Category < ActiveRecord::Base
  validates :name, presence: true, uniqueness: { case_sensitive: false }

  scope :income, -> { where('income IN(?)', [1, true]) }

  scope :expense, -> { where('income IN(?)', [0, false]) }

  before_save :assign_slug

  acts_as_paranoid

  private
  #
  # TODO: BUG
  #
  def assign_slug
    #self.slug = Russian::Transliteration.transliterate(self.name).downcase.gsub(/[^a-z]+/, '_')
  end

  class << self
    def group_by_income
      [
        ['Видатки', expense.pluck(:name, :id)],
        ['Надходження', income.pluck(:name, :id)]
      ]
    end
  end
end
