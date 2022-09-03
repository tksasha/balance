# frozen_string_literal: true

RSpec.describe ConsolidationExpensesSum do
  its(:name) { is_expected.to eq 'Сума витрат' }

  its(:sum) { is_expected.to eq 0 }

  its(:income?) { is_expected.to be false }

  its(:category_id) { is_expected.to be_nil }

  context do
    before { described_class.sum = 42.69 }

    after { described_class.sum = 0 }

    its(:sum) { is_expected.to eq 42.69 }
  end
end
