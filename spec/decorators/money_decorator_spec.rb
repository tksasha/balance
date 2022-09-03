# frozen_string_literal: true

RSpec.describe MoneyDecorator do
  subject { described_class.new money }

  let(:money) { 241_445 }

  its(:to_s) { is_expected.to eq '241 445.00' }
end
