# frozen_string_literal: true

RSpec.describe MoneyHelper do
  subject { helper }

  describe '#money' do
    subject { helper.money 400_500.2 }

    it { is_expected.to eq '400 500.20' }
  end
end
