class Category < ActiveRecord::Base
  validates :name, presence: true, uniqueness: { case_sensitive: false }

  scope :income, -> { where('income IN(?)', [1, true]) }

  scope :expense, -> { where('income IN(?)', [0, false]) }

  scope :visible, -> { where visible: true }

  before_save :assign_slug

  def destroy
    update visible: false
  end

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
        ['Видатки', visible.expense.pluck(:name, :id)],
        ['Надходження', visible.income.pluck(:name, :id)]
      ]
    end
  end
end
