# frozen_string_literal: true

class Category < ApplicationRecord
  validates :name, presence: true, uniqueness: { case_sensitive: false, scope: :currency }

  validates :currency, presence: true

  enum currency: CURRENCIES

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }

  before_save :assign_slug

  private

  def assign_slug
    self.slug = SlugService.build name
  end
end
