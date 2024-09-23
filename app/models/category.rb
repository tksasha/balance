# frozen_string_literal: true

# == Schema Information
#
# Table name: categories
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  deleted_at    :datetime
#  income        :boolean          default(FALSE)
#  name          :string
#  supercategory :integer          default("common"), not null
#  visible       :boolean          default(TRUE)
#
# Indexes
#
#  index_categories_on_name_and_currency  (name,currency) UNIQUE
#

class Category < ApplicationRecord
  validates :name, presence: true, uniqueness: { case_sensitive: false, scope: :currency }

  validates :currency, presence: true

  enum :currency, CURRENCIES

  enum \
    :supercategory,
    { common: 1, children: 2, business: 3, charity: 4, currency: 5 },
    default: :common, scopes: false

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }

  default_scope -> { where(deleted_at: nil) }

  def destroy
    touch(:deleted_at)
  end

  class << self
    # TODO: spec me
    def ransackable_attributes(*)
      %w[currency id income name supercategory visible]
    end
  end
end
