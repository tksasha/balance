# frozen_string_literal: true

# == Schema Information
#
# Table name: cashes
#
#  id         :integer          not null, primary key
#  currency   :integer          default("uah")
#  deleted_at :time
#  formula    :string
#  name       :string
#  sum        :decimal(10, 2)
#
# Indexes
#
#  index_cashes_on_name_and_currency  (name,currency) UNIQUE
#
class Cash < ApplicationRecord
  include ActsAsHasFormula

  validates :name, :formula, :currency, presence: true

  validates :name, uniqueness: { case_sensitive: false, scope: :currency }

  enum currency: CURRENCIES

  has_paper_trail
end
