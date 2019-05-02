require 'rails_helper'

RSpec.describe DateFactory do
  describe '.build' do
    before { allow(Date).to receive(:today) { Date.new(2013, 3, 31) } }

    context do
      subject { DateFactory.build }

      it { should eq Date.new 2013, 3, 1 }
    end

    context do
      subject { DateFactory.build year: '2011' }

      it { expect(subject.year).to eq 2011 }
    end

    context do
      subject { DateFactory.build month: '11' }

      it { expect(subject.month).to eq 11 }
    end

    context do
      subject { DateFactory.build day: '17' }

      it { expect(subject.day).to eq 17 }
    end

    context do
      subject { DateFactory.build year: 2013, month: 2 }

      it { should eq Date.new 2013, 2, 1 }
    end
  end
end
