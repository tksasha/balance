# frozen_string_literal: true

RSpec.describe MoneyHelper, type: :helper do
  subject { helper }

  describe '#money' do
    subject { helper.money 400_500.2 }

    it { should eq '400 500.20' }
  end
end
