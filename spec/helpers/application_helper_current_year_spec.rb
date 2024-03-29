# frozen_string_literal: true

RSpec.describe ApplicationHelper do
  subject { helper }

  describe '#current_year' do
    let(:month) { Month.new(2021, 3) }

    before { allow(subject).to receive(:month).and_return(month) }

    its(:current_year) { is_expected.to eq 2021 }
  end
end
